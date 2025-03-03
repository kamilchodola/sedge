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
package services

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (s *serviceManager) ContainerId(service string) (string, error) {
	containers, err := s.dockerClient.ContainerList(context.Background(), types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", service)),
	})
	if err != nil {
		return "", err
	}
	if len(containers) == 0 {
		return "", fmt.Errorf("%w: %s", ErrContainerNotFound, service)
	}
	if len(containers) > 1 {
		return "", fmt.Errorf("%w: %s", ErrMultipleContainers, service)
	}
	return containers[0].ID, nil
}
