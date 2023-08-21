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
	"bytes"
	"go.resf.org/peridot/devtools/garrison/internal/pkg/starlark"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (s *State) NewService(service *starlark.Service) error {
	md := s.getMetadata(service)

	podLabels := map[string]string{
		"app":   *service.Name,
		"stage": string(s.Stage),
	}

	// Copy the metadata and add the labels.
	podMd := *md
	podMd.Labels = podLabels

	// Copy the metadata and add the labels.
	svcMd := *md
	svcMd.Labels = podLabels

	// Create the service.
	svc := &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: svcMd,
		Spec: v1.ServiceSpec{
			Selector: podLabels,
			Ports:    []v1.ServicePort{},
		},
	}

	// Create the ports.
	for _, port := range service.Ports {
		svc.Spec.Ports = append(svc.Spec.Ports, v1.ServicePort{
			Name:       *port.Name,
			Port:       int32(*port.Port),
			TargetPort: intstr.FromInt32(int32(*port.Port)),
		})
	}

	// Create yaml buffer
	var outBuffer bytes.Buffer
	err := s.printer.PrintObj(svc, &outBuffer)
	if err != nil {
		return err
	}

	s.manifests = append(s.manifests, outBuffer.String())

	return nil
}
