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
