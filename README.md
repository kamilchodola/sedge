# Sedge

[![Go Report Card](https://goreportcard.com/badge/github.com/NethermindEth/sedge)](https://goreportcard.com/report/github.com/NethermindEth/sedge)
[![Discord](https://user-images.githubusercontent.com/7288322/34471967-1df7808a-efbb-11e7-9088-ed0b04151291.png)](https://discord.com/invite/PaCMRFdvWT)
[![codecov](https://codecov.io/gh/NethermindEth/sedge/branch/main/graph/badge.svg?token=8FERO4PO1V)](https://codecov.io/gh/NethermindEth/sedge)

A one-click setup tool for PoS network/chain validators. Currently, Sedge is designed primarily for solo stakers and testnet devs of Ethereum. Sedge generates docker-compose scripts for the entire on-premise validator setup based on the chosen client.

The project **is still on beta** and although it should be stable enough, it still might have some issues. Sedge has not been audited yet.

- [Sedge](#sedge)
  - [⚙️ Installation](#️-installation)
    - [Dependencies](#dependencies)
    - [Installation methods](#installation-methods)
    - [**Disclaimer**](#disclaimer)
  - [📜 Documentation](#-documentation)
  - [⚡️ Quick start](#️-quick-start)
  - [💥 Why did we start Sedge?](#-why-did-we-start-sedge)
  - [🔥 What can you do with Sedge today?](#-what-can-you-do-with-sedge-today)
    - [**Disclaimer**](#disclaimer-1)
  - [Supported networks and clients](#supported-networks-and-clients)
    - [Mainnet](#mainnet)
    - [Sepolia](#sepolia)
    - [Goerli](#goerli)
    - [Gnosis](#gnosis)
    - [Chiado (Gnosis testnet)](#chiado-gnosis-testnet)
    - [CL clients with Mev-Boost](#cl-clients-with-mev-boost)
  - [Supported Linux flavours for dependency installation](#supported-linux-flavours-for-dependency-installation)
  - [✅ Roadmap](#-roadmap)
    - [Version 0.1](#version-01)
    - [Version 0.2](#version-02)
    - [Version 0.3](#version-03)
    - [Version 0.4](#version-04)
    - [Version 0.5](#version-05)
    - [Version 0.6 (Actual)](#version-06-actual)
    - [Version 0.X](#version-0x)
    - [Version 1.0](#version-10)
  - [Version X.0](#version-x0)
  - [💪 Want to contribute?](#-want-to-contribute)
  - [⚠️ License](#️-license)

## ⚙️ Installation

### Dependencies

Sedge dependencies are `docker` with `docker compose` plugin, but if you don't have those installed, Sedge will show instructions to install them, or install them for you. Check the [docs](https://docs.sedge.nethermind.io/docs/quickstart/dependencies) for more details.

### Installation methods

Check our [installation guide](https://docs.sedge.nethermind.io/docs/quickstart/install-guide) for detailed instructions on the supported methods:

- Download binary from release page
- Using the Homebrew package manager
- Using the Go programming language
- Build from source

### **Disclaimer**

Downloading any binary from the internet comes with the risk of downloading files that malicious, third-party actors have injected with malware. All users should check that they are downloading the correct, clean binary from a reputable source.

## 📜 Documentation

For further details, you can check the [documentation](https://docs.sedge.nethermind.io).

## ⚡️ Quick start

With `sedge cli` you can go through the entire workflow setup:

1. Check dependencies
2. Generate jwtsecret (not for mainnet)
3. Generate a `docker-compose` script with randomized clients selection and `.env`
4. Execute the `docker-compose` script (only execution and consensus nodes will be executed by default)
5. Validator client will be executed automatically after execution and consensus nodes are synced.
  
Between steps 4 and 5 you can generate the validator(s) keystore folder using `sedge keys`.

The entire process is interactive, although you can use the `-y` flag to run Sedge without prompts.

Check all the options and flags with `sedge cli --help`. More instructions or guides about Sedge's features are coming soon!

## 💥 Why did we start Sedge?

As people who actively deployed validators way before The Merge, we know how hard it is to set up an Ethereum validator:

- You need to procure at least three (compatible) nodes: an execution node (geth, nethermind, erigon, etc), a consensus node, and a validator node (lighthouse, prysm, etc)
- You then need to execute them, connect them, monitor them, and secure the validator keys (which includes staking 32 ETH).
- There may be several valid combinations of clients to choose for your setup, so you need to go through each of the client's docs, evaluate it, get instructions for it and test it. You also need to feel comfortable executing commands in the cli, know docker, and understand basics of networking. On top of this, there are many different settings you must read up on and consider for your client node.
- In the case of working with the Ethereum Mainnet, you are working with real money that can potentially be lost in the event of having downtime or being slashed. To avoid losing real value, you must be aware of and follow best practices on the validator setup, and correctly monitor your nodes.
- Have you heard of MEV? Flashbots is working on an MEV-Boost component which will take your validator to another level of awesomeness. You most likely want to always be running the latest version, but you also most likely don’t have the time to understand the MEV-Boost architecture in and out, or how to successfully implement it into your environment.
  
> Enter sedge

We want Sedge to take care of all of the above for you. With just a few clicks or steps, Sedge can create an entire Ethereum staking architecture that supports client diversity and Ethereum's latest features, while being completely free and open source. We want Sedge to save you from making costly mistakes in this complex setup; along with hours or days of research, reading and testing. We want you to be able to stake easily with or without blockchain knowledge by giving you the tools to help this amazing community (and earn some good money of course 😉).

We want to share our knowledge in this topic and create something that allows everyone to easily and safely set up lots of diverse validators.

We don't want to stop at Ethereum. We also want to help stakers of other PoS networks/chains, so if your favourite chain is not here, you are more than welcome to contribute!

## 🔥 What can you do with Sedge today?

- Select an execution, consensus and validator client node (manually or automatically) and generate a `docker-compose` script with production-tested configurations to run the setup you want.
- Don't remember `docker-compose` commands or flags for your setup? Check docker logs of the running services with `sedge logs`, and shut them down with `sedge down`
- Generate the keystore folder with `sedge keys` for mainnet using the [staking-deposit-cli](https://github.com/ethereum/staking-deposit-cli) tool or for the rest of supported networks using our own experimental code.

> **Disclaimer:** Users acknowledge that staking-deposit-cli is an external tool, which means that Nethermind exercises no control over its functioning and does not accept any liability for any issues that may arise from the use of the tool.

> **Disclaimer:** Users acknowledge that generating the keystore for any network other than the mainnet is an experimental and unaudited feature of Sedge. Nethermind provides this feature on an ‘as is’ basis and makes no warranties regarding its proper functioning. The use of this feature is at the user’s own risk - Nethermind excludes all liability for any malfunction or loss of money that may occur as the result of an unexpected behavior during the keystore generation.

The setup is currently designed to start all three nodes required to run a full local validator (execution, consensus and validator node). Soon, Sedge will let you connect to a public or remote node for the execution and consensus layers. Once the consensus node is synced, a validator node will be executed automatically. We suggest that you make use of the initial time taken to sync to prepare the keystore file and make the deposit for your staked ether.

If you are familiar with `docker`, `docker compose`, and the validator setup, then you can use Sedge to generate a base docker-compose script with the recommended settings, stop Sedge instead of letting it execute the script, and then edit the script as much as you want. It is a lot more easier than doing everything from scratch!

> Although Sedge supports several clients, **is still on beta**. Some settings may not work because -at least on the testnets- the clients are constantly evolving. Please let us know about any issues you encounter!

### **Disclaimer**

While Sedge assists in installing the validator, it is not designed to register or maintain it. Users are solely responsible for ensuring that they monitor and maintain the validator as required, so that they do not incur penalties and/or financial losses. This includes promptly updating the tool to ensure the latest stable releases of clients are used.

Users acknowledge that no warranty is being made of a successful installation. Sedge is a tool and it ultimately depends on you to use it correctly and follow all the best practice guidance, as found in this README and documentation.

## Supported networks and clients

### Mainnet

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Geth       | Lighthouse | Lighthouse |
| Nethermind | Lodestar   | Lodestar   |
| Erigon     | Prysm      | Prysm      |
| Besu       | Teku       | Teku       |

### Sepolia

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Geth       | Lighthouse | Lighthouse |
| Nethermind | Lodestar   | Lodestar   |
| Erigon     | Prysm      | Prysm      |
| Besu       | Teku       | Teku       |

### Goerli

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Geth       | Lighthouse | Lighthouse |
| Nethermind | Lodestar   | Lodestar   |
| Erigon     | Prysm      | Prysm      |
| Besu       | Teku       | Teku       |

### Gnosis

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Nethermind | Lighthouse | Lighthouse |
|            | Lodestar   | Lodestar   |
|            | Teku       | Teku       |

### Chiado (Gnosis testnet)

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Nethermind | Lighthouse | Lighthouse |
|            | Lodestar   | Lodestar   |
|            | Teku       | Teku       |

### CL clients with Mev-Boost

| Client     | Mev-Boost | Networks                 |
| ---------- | --------- |--------------------------|
| Lighthouse | yes       | Mainnet, Goerli, Sepolia |
| Lodestar   | yes       | Mainnet, Goerli, Sepolia |
| Prysm      | yes       | Mainnet, Goerli, Sepolia |
| Teku       | yes       | Mainnet, Goerli, Sepolia |

## Supported Linux flavours for dependency installation

| OS             | Versions                |
| -------------- | ----------------------- |
| Ubuntu         | 22.04, 20.04 |
| Debian         | 11,10,9,8               |
| Fedora         | 35,34                   |
| CentOS         | 8                       |
| Arch           | -                       |
| Amazon Linux 2 | -                       |
| Alpine         | 3.15,3.14,3.14.3        |

## ✅ Roadmap

The following roadmap covers the main features and ideas we want to implement but doesn't cover everything we are planning for this tool. Stay in touch if you are interested, a lot of improvements are coming in the next two months. Please note that this Roadmap is continually changing until version 1.0.

### Version 0.1

- [x] Generate `docker-compose` scripts and `.env` files for selected clients with a cli tool
- [x] Generate keystore folder with the cli
- [x] Test coverage (unit tests)
- [x] Integrate Kiln network
- [x] Integrate MEV-Boost as an option
- [x] Integrate Ropsten network

### Version 0.2

- [x] Integrate Goerli/Prater network
- [x] Integrate Sepolia network
- [x] Documentation with examples

### Version 0.3

- [x] Integrate Gnosis network
- [x] Prepare for the Merge

### Version 0.4

- [x] Create and handle keystores on our own instead of using staking-deposit-cli
- [x] Improve validator testing
- [x] Bug fixes
- [x] Deprecate Kiln, Ropsten, Denver networks
- [x] Improve support for chiado network (Gnosis testnet)

### Version 0.5

- [x] Support for Gnosis Merge
- [x] Bug fixes

### Version 0.6 (Actual)

- [x] Besu and Erigon support
- [x] Windows support
- [x] Bug fixes

### Version 0.X

- [ ] Set up and run only one node (execution/consensus/validator)
- [ ] Include monitoring tool for alerting, tracking validator balance, and tracking sync progress and status of nodes
- [ ] More tests!!!
- [ ] Support for Nimbus client

### Version 1.0

Full Ethereum PoS support with MEV-Boost

## Version X.0

- [ ] Integrate other PoS networks
- [ ] TUI for guided and more interactive setup (better UX)
- [ ] Off-premise setup support

## 💪 Want to contribute?

Please check our Contributing Guidelines, Code of Conduct and our issues. In case you want to report or suggest something (any help is welcome), please file an issue first so that the main team is aware and can discuss it.

If you know of any good tricks for validator setup that other people could make good use of as well, please consider adding it to Sedge. Your efforts will be greatly appreciated by the community.

## ⚠️ License

Sedge is a Nethermind free and open-source software licensed under the [Apache 2.0 License](https://github.com/NethermindEth/sedge/blob/main/LICENSE).
