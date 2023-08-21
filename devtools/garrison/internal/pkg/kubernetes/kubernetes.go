// Copyright 2023 Peridot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubernetes

import (
	"go.resf.org/peridot/devtools/garrison/internal/pkg/starlark"
	"k8s.io/cli-runtime/pkg/printers"
	"strings"
)

type State struct {
	*starlark.State

	printer   *printers.YAMLPrinter
	manifests []string
	helmMode  bool
}

func NewState(state *starlark.State) *State {
	return &State{
		State:     state,
		printer:   &printers.YAMLPrinter{},
		manifests: []string{},
		helmMode:  false,
	}
}

func (s *State) GenerateManifests() error {
	for _, service := range s.Services.Services {
		err := s.NewDeployment(service)
		if err != nil {
			return err
		}

		err = s.NewService(service)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *State) YAML() (string, error) {
	var outBuffer strings.Builder
	outBuffer.WriteString("---\n")

	for _, manifest := range s.manifests {
		outBuffer.WriteString(manifest)
	}

	return outBuffer.String(), nil
}
