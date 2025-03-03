/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package generate

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

const (
	wrongDep = "wrong_dep"
)

type CheckFunc func(t *testing.T, data *GenData, compose, env io.Reader) error

type genTestData struct {
	Description     string
	GenerationData  *GenData
	ErrorGenCompose error
	ErrorGenEnvFile error
	ErrorCheckFunc  error
	CheckFunctions  []CheckFunc
}

func clean(s string) string {
	return strings.ReplaceAll(s, "\r", "")
}

var checkOnlyExecution = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	composeData, err := retrieveComposeData(compose)
	if err != nil {
		return err
	}
	assert.NotNil(t, composeData.Services)
	assert.NotNil(t, composeData.Services.Execution)
	assert.Equal(t, composeData.Services.Execution.ContainerName, "execution-client")
	return nil
}

var checkOnlyConsensus = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	composeData, err := retrieveComposeData(compose)
	if err != nil {
		return err
	}
	assert.NotNil(t, composeData.Services)
	assert.NotNil(t, composeData.Services.Consensus)
	assert.Equal(t, composeData.Services.Consensus.ContainerName, "consensus-client")
	return nil
}

var checkOnlyValidator = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	composeData, err := retrieveComposeData(compose)
	if err != nil {
		return err
	}
	assert.NotNil(t, composeData.Services)
	assert.NotNil(t, composeData.Services.Validator)
	assert.Equal(t, composeData.Services.Validator.ContainerName, "validator-client")
	return nil
}

var checkCCBootnodesOnConsensus = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	composeData, err := retrieveComposeData(compose)
	if err != nil {
		return err
	}
	var ccBootnodes []string
	if data.CCBootnodes != nil {
		ccBootnodes = *data.CCBootnodes
	}
	if len(ccBootnodes) == 0 {
		ccBootnodes = configs.NetworksConfigs()[data.Network].DefaultCCBootnodes
	}
	if len(ccBootnodes) != 0 {
		bootnodes := strings.Join(ccBootnodes, ",")
		if composeData.Services.Consensus != nil && data.ConsensusClient.Name == "lighthouse" {
			checkFlagOnCommands(t, composeData.Services.Consensus.Command, "--boot-nodes="+bootnodes)
		}
		if composeData.Services.Consensus != nil && data.ConsensusClient.Name == "prysm" {
			for _, bNode := range ccBootnodes {
				checkFlagOnCommands(t, composeData.Services.Consensus.Command, "--bootstrap-node="+bNode)
			}
		}
		if composeData.Services.Consensus != nil && data.ConsensusClient.Name == "teku" {
			checkFlagOnCommands(t, composeData.Services.Consensus.Command, "--p2p-discovery-bootnodes="+bootnodes)
		}
		if composeData.Services.Consensus != nil && data.ConsensusClient.Name == "lodestar" {
			for _, bNode := range ccBootnodes {
				checkFlagOnCommands(t, composeData.Services.Consensus.Command, "--bootnodes="+bNode)
			}
		}
	}
	return nil
}

var checkTTDOnExecution = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	composeData, err := retrieveComposeData(compose)
	if err != nil {
		return err
	}
	customTTD := data.CustomTTD
	if customTTD == "" {
		customTTD = configs.NetworksConfigs()[data.Network].DefaultTTD
	}
	if customTTD != "" {
		if composeData.Services.Execution != nil && data.ExecutionClient.Name == "besu" {
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--override-genesis-config=terminalTotalDifficulty="+customTTD)
		}
		if composeData.Services.Execution != nil && data.ExecutionClient.Name == "erigon" {
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--override.terminaltotaldifficulty="+customTTD)
		}
		if composeData.Services.Execution != nil && data.ExecutionClient.Name == "nethermind" {
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--Merge.TerminalTotalDifficulty="+customTTD)
		}
		if composeData.Services.Execution != nil && data.ExecutionClient.Name == "geth" {
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--override.terminaltotaldifficulty="+customTTD)
		}
	}
	return nil
}

var checkECBootnodesOnExecution = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	composeData, err := retrieveComposeData(compose)
	if err != nil {
		return err
	}
	var ecBootnodes []string
	if data.ECBootnodes != nil {
		ecBootnodes = *data.ECBootnodes
	}
	if len(ecBootnodes) == 0 {
		ecBootnodes = configs.NetworksConfigs()[data.Network].DefaultECBootnodes
	}
	if len(ecBootnodes) != 0 {
		bootnodes := strings.Join(ecBootnodes, ",")
		if composeData.Services.Execution != nil && data.ExecutionClient.Name == "besu" {
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--bootnodes="+bootnodes)
		}
		if composeData.Services.Execution != nil && data.ExecutionClient.Name == "erigon" {
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--bootnodes="+bootnodes)
		}
		if composeData.Services.Execution != nil && data.ExecutionClient.Name == "nethermind" {
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--Network.Bootnodes="+bootnodes)
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--Discovery.Bootnodes="+bootnodes)
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--Network.StaticPeers="+bootnodes)
		}
		if composeData.Services.Execution != nil && data.ExecutionClient.Name == "geth" {
			checkFlagOnCommands(t, composeData.Services.Execution.Command, "--bootnodes="+bootnodes)
		}
	}
	return nil
}

// retrieveComposeData returns compose data from the reader
func retrieveComposeData(compose io.Reader) (*ComposeData, error) {
	// load compose file
	composeBytes, err := io.ReadAll(compose)
	if err != nil {
		return nil, err
	}
	var composeData ComposeData
	err = yaml.Unmarshal(composeBytes, &composeData)
	if err != nil {
		return nil, err
	}
	return &composeData, nil
}

func checkFlagOnCommands(t *testing.T, commands []string, flag string) {
	t.Helper()
	exists := false
	for _, line := range commands {
		if strings.Contains(line, flag) {
			exists = true
			break
		}
	}
	assert.True(t, exists)
}

var checkMevServices = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	// load compose file
	composeBytes, err := io.ReadAll(compose)
	if err != nil {
		return err
	}
	var composeData ComposeData
	err = yaml.Unmarshal(composeBytes, &composeData)
	if err != nil {
		return err
	}

	if utils.Contains(data.Services, mevBoost) {
		if composeData.Services.Mevboost != nil {
			assert.Equal(t, "flashbots/mev-boost:latest", composeData.Services.Mevboost.Image)
			assert.Equal(t, "mev-boost", composeData.Services.Mevboost.ContainerName)
			assert.Equal(t, "on-failure", composeData.Services.Mevboost.Restart)
		} else {
			return errors.New("mevboost service is not present")
		}
	}

	return nil
}

var checkExtraFlagsOnExecution = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	// load compose file
	composeBytes, err := io.ReadAll(compose)
	if err != nil {
		return err
	}
	var composeData ComposeData
	err = yaml.Unmarshal(composeBytes, &composeData)
	if err != nil {
		return err
	}

	if composeData.Services.Execution != nil {
		var extraFlags []string
		if data.ElExtraFlags != nil {
			extraFlags = *data.ElExtraFlags
		}
		for _, flag := range extraFlags {
			assert.True(t, utils.Contains(composeData.Services.Execution.Command, "--"+flag))
		}
	} else {
		return errors.New("execution service is not present")
	}

	return nil
}

var defaultFunc = func(t *testing.T, data *GenData, compose, env io.Reader) error {
	// load compose file
	composeBytes, err := io.ReadAll(compose)
	if err != nil {
		return err
	}
	var composeData ComposeData
	err = yaml.Unmarshal(composeBytes, &composeData)
	if err != nil {
		return err
	}

	if utils.Contains(data.Services, execution) {
		assert.NotNil(t, composeData.Services.Execution)
	}
	if utils.Contains(data.Services, consensus) {
		assert.NotNil(t, composeData.Services.Consensus)
	}
	if utils.Contains(data.Services, validator) {
		assert.NotNil(t, composeData.Services.Validator)
	}
	if utils.Contains(data.Services, mevBoost) {
		assert.NotNil(t, composeData.Services.Mevboost)
	}
	if utils.Contains(data.Services, validatorImport) {
		assert.NotNil(t, composeData.Services.ValidatorImport)
	}
	if utils.Contains(data.Services, configConsensus) {
		assert.NotNil(t, composeData.Services.ConfigConsensus)
	}

	// load .env file
	envData := retrieveEnvData(t, env)
	if data.Network == "gnosis" {
		// Check that the right network is set
		assert.Contains(t, envData, "EL_NETWORK")
		assert.Equal(t, "xdai", clean(envData["EL_NETWORK"]))

		assert.Contains(t, envData, "CL_NETWORK")
		assert.Equal(t, "gnosis", clean(envData["CL_NETWORK"]))

	} else {
		// Check that the right network is set
		assert.Contains(t, envData, "NETWORK")
		assert.Equal(t, data.Network, clean(envData["NETWORK"]))
	}

	return nil
}

func generateTestCases(t *testing.T) (tests []genTestData) {
	baseDescription := "Test generation of compose services "
	tests = []genTestData{{Description: baseDescription + " NilData", ErrorGenCompose: ErrEmptyData, CheckFunctions: []CheckFunc{defaultFunc}}}

	networks, err := utils.SupportedNetworks()
	if err != nil {
		t.Error("SupportedNetworks() failed", err)
	}

	for _, network := range networks {
		c := clients.ClientInfo{Network: network}

		executionClients, err := c.SupportedClients("execution")
		if err != nil {
			t.Errorf("SupportedClients(\"execution\") failed: %v", err)
		}
		consensusClients, err := c.SupportedClients("consensus")
		if err != nil {
			t.Errorf("SupportedClients(\"consensus\") failed: %v", err)
		}
		validatorClients, err := c.SupportedClients("validator")
		if err != nil {
			t.Errorf("SupportedClients(\"validator\") failed: %v", err)
		}

		// TODO: Add CheckpointSyncUrl, FallbackELUrls and FeeRecipient to test data
		for _, executionCl := range executionClients {
			for _, consensusCl := range consensusClients {
				tests = append(tests,
					genTestData{
						Description: fmt.Sprintf(baseDescription+"execution: %s, network: %s, only execution", executionCl, network),
						GenerationData: &GenData{
							ExecutionClient: &clients.Client{Name: executionCl},
							Network:         network,
							Services:        []string{execution},
						},
						CheckFunctions: []CheckFunc{defaultFunc, checkOnlyExecution},
					},
					genTestData{
						Description: fmt.Sprintf(baseDescription+"consensus: %s, network: %s, only consensus", consensusCl, network),
						GenerationData: &GenData{
							ConsensusClient: &clients.Client{Name: consensusCl},
							Network:         network,
							Services:        []string{consensus},
						},
						CheckFunctions: []CheckFunc{defaultFunc, checkOnlyConsensus},
					},
					genTestData{
						Description: fmt.Sprintf(baseDescription+"validator: %s, network: %s, only validator", consensusCl, network),
						GenerationData: &GenData{
							ValidatorClient: &clients.Client{Name: consensusCl},
							Network:         network,
							Services:        []string{validator},
						},
						CheckFunctions: []CheckFunc{defaultFunc, checkOnlyValidator},
					},
				)
				if utils.Contains(validatorClients, consensusCl) {
					tests = append(tests,
						genTestData{
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s, all", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: consensusCl},
								ValidatorClient: &clients.Client{Name: consensusCl},
								Network:         network,
								Services:        []string{execution, consensus, validator},
							},
							CheckFunctions: []CheckFunc{defaultFunc},
						},
						genTestData{
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s, no validator", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								Services:        []string{execution, consensus},
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: consensusCl},
								Network:         network,
							},
							CheckFunctions: []CheckFunc{defaultFunc},
						},
						genTestData{
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s, Execution Client not Valid", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								Services:        []string{execution, consensus},
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: wrongDep},
								Network:         network,
							},
							ErrorGenCompose: ErrConsensusClientNotValid,
							CheckFunctions:  []CheckFunc{defaultFunc},
						})
				}
			}
		}
	}

	return
}

func TestGenerateComposeServices(t *testing.T) {
	configs.InitNetworksConfigs()
	tests := []genTestData{
		{
			Description: "Test generation of compose services",
			GenerationData: &GenData{
				Services:        []string{execution, consensus, validator, validatorImport, mevBoost},
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			CheckFunctions: []CheckFunc{defaultFunc},
		},
		{
			Description: "Test mevBoost service",
			GenerationData: &GenData{
				Services:        []string{mevBoost},
				Network:         "mainnet",
				Mev:             true,
				MevBoostService: true,
			},
			CheckFunctions: []CheckFunc{defaultFunc, checkMevServices},
		},
		{
			Description: "Test EL extra flags",
			GenerationData: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				Services:        []string{execution},
				Network:         "mainnet",
				ElExtraFlags:    &[]string{"extra", "flag"},
			},
			CheckFunctions: []CheckFunc{checkExtraFlagsOnExecution},
		},
	}

	tests = append(tests, generateTestCases(t)...)

	tests = append(tests, customFlagsTestCases(t)...)

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			var buffer bytes.Buffer
			err := ComposeFile(tt.GenerationData, io.Writer(&buffer))
			assert.ErrorIs(t, err, tt.ErrorGenCompose)
			if tt.ErrorGenCompose != nil {
				return
			}

			var envBuffer bytes.Buffer
			err = EnvFile(tt.GenerationData, io.Writer(&envBuffer))
			assert.ErrorIs(t, err, tt.ErrorGenEnvFile)
			if tt.ErrorGenEnvFile != nil {
				return
			}

			for _, f := range tt.CheckFunctions {
				err = f(t, tt.GenerationData, bytes.NewReader(buffer.Bytes()), bytes.NewReader(envBuffer.Bytes()))
				assert.ErrorIs(t, err, tt.ErrorCheckFunc)

			}
		})
	}
}

func customFlagsTestCases(t *testing.T) (tests []genTestData) {
	baseDescription := "Test generation of compose services "
	tests = []genTestData{}

	for _, network := range []string{"custom", "mainnet"} {
		c := clients.ClientInfo{Network: network}

		executionClients, err := c.SupportedClients("execution")
		if err != nil {
			t.Errorf("SupportedClients(\"execution\") failed: %v", err)
		}
		consensusClients, err := c.SupportedClients("consensus")
		if err != nil {
			t.Errorf("SupportedClients(\"consensus\") failed: %v", err)
		}
		validatorClients, err := c.SupportedClients("validator")
		if err != nil {
			t.Errorf("SupportedClients(\"validator\") failed: %v", err)
		}

		for _, executionCl := range executionClients {
			for _, consensusCl := range consensusClients {
				if utils.Contains(validatorClients, consensusCl) {
					tests = append(tests,
						genTestData{
							Description: fmt.Sprintf(baseDescription+"customTTD tests, execution: %s, consensus: %s, validator: %s, network: %s, all", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: consensusCl},
								ValidatorClient: &clients.Client{Name: consensusCl},
								Network:         network,
								CustomTTD:       "sample_ttd",
								Services:        []string{execution, consensus, validator},
							},
							CheckFunctions: []CheckFunc{checkTTDOnExecution, defaultFunc, checkECBootnodesOnExecution},
						},
						genTestData{
							Description: fmt.Sprintf(baseDescription+"ecBootnodes tests, execution: %s, consensus: %s, validator: %s, network: %s, no validator", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								Services:        []string{execution, consensus},
								ExecutionClient: &clients.Client{Name: executionCl},
								ECBootnodes:     &[]string{"enode:1", "enode:2", "enode:3"},
								ConsensusClient: &clients.Client{Name: consensusCl},
								Network:         network,
							},
							CheckFunctions: []CheckFunc{defaultFunc, checkECBootnodesOnExecution, checkTTDOnExecution},
						},
						genTestData{
							Description: fmt.Sprintf(baseDescription+"ccBootnodes tests, execution: %s, consensus: %s, validator: %s, network: %s, Execution Client not Valid", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								Services:        []string{consensus},
								ConsensusClient: &clients.Client{Name: consensusCl},
								CCBootnodes:     &[]string{"enr:1", "enr:2"},
								Network:         network,
							},
							CheckFunctions: []CheckFunc{defaultFunc, checkCCBootnodesOnConsensus},
						})
				}
			}
		}
	}

	return
}

// TestValidateClients tests the validation of clients
func TestValidateClients(t *testing.T) {
	configs.InitNetworksConfigs()
	tests := []struct {
		Description string
		Data        *GenData
		Error       error
	}{
		{
			Description: "Wrong execution client",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "wrong"},
				Network:         "mainnet",
			},
			Error: ErrExecutionClientNotValid,
		},
		{
			Description: "Wrong consensus client",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "wrong"},
				Network:         "mainnet",
			},
			Error: ErrConsensusClientNotValid,
		},
		{
			Description: "Wrong validator client",
			Data: &GenData{
				ValidatorClient: &clients.Client{Name: "wrong"},
				Network:         "mainnet",
			},
			Error: ErrValidatorClientNotValid,
		},
		{
			Description: "Wrong network, empty clients",
			Data: &GenData{
				Network: wrongDep,
			},
			Error: nil,
		},
		{
			Description: "Wrong network, good consensus",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "teku"},
				Network:         wrongDep,
			},
			Error: ErrUnableToGetClientsInfo,
		},
		{
			Description: "Wrong network, good execution",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				Network:         wrongDep,
			},
			Error: ErrUnableToGetClientsInfo,
		},
		{
			Description: "Wrong network, good validator",
			Data: &GenData{
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         wrongDep,
			},
			Error: ErrUnableToGetClientsInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			err := validateClients(tt.Data)
			assert.ErrorIs(t, err, tt.Error)
		})
	}
}

func TestEnvFileAndFlags(t *testing.T) {
	configs.InitNetworksConfigs()
	tests := []struct {
		Description string
		Data        *GenData
		Error       error
	}{
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Error: nil,
		},
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku", Endpoint: "http://localhost"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Error: nil,
		},
		{
			Description: "Empty data",
			Data: &GenData{
				Network: "mainnet",
			},
			Error: ErrEmptyData,
		},
		{
			Description: "Wrong network",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "teku"},
				Network:         wrongDep,
			},
			Error: ErrTemplateNotFound,
		},
		{
			Description: "Prysm consensus with ConsensusAdditionalUrl",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "prysm"},
				ValidatorClient: &clients.Client{Name: "prysm"},
				Network:         "mainnet",
				ConsensusApiUrl: "http://localhost:8080",
			},
			Error: nil,
		},
		{
			Description: "Teku consensus and validator with ConsensusApiUrl",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				ConsensusApiUrl: "http://localhost:8080",
			},
			Error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			var buffer bytes.Buffer
			err := EnvFile(tt.Data, io.Writer(&buffer))
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
				return
			}
			if tt.Data.ConsensusClient != nil {
				if tt.Data.ConsensusApiUrl == "" {
					str := buffer.String()
					assert.Contains(t, str, "CC_API_URL="+endpointOrEmpty(tt.Data.ConsensusClient)+":")
				} else {
					if tt.Data.ConsensusClient.Name == "prysm" && tt.Data.ValidatorClient != nil {
						assert.Contains(t, buffer.String(), "CC_API_URL=consensus:")
					} else {
						assert.Contains(t, buffer.String(), "CC_API_URL="+tt.Data.ConsensusApiUrl)
					}
				}
			}
		})
	}
}

func TestCleanGeneratedFiles(t *testing.T) {
	configs.InitNetworksConfigs()
	tests := []struct {
		Description string
		Data        *GenData
	}{
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
		},
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			path := t.TempDir()

			// generate files
			out, err := os.Create(filepath.Join(path, configs.DefaultDockerComposeScriptName))
			if err != nil {
				return
			}
			defer out.Close()
			assert.Nil(t, err)
			err = ComposeFile(tt.Data, out)
			assert.Nil(t, err)
			assert.FileExists(t, filepath.Join(path, configs.DefaultDockerComposeScriptName))

			// open output file
			out, err = os.Create(filepath.Join(path, configs.DefaultEnvFileName))
			assert.ErrorIs(t, err, nil)
			defer out.Close()
			err = EnvFile(tt.Data, out)
			assert.Nil(t, err)
			assert.FileExists(t, filepath.Join(path, configs.DefaultEnvFileName))

			err = CleanGenerated(path)
			assert.Nil(t, err)
		})
	}
}
