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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *State) getMetadata(service *starlark.Service) *metav1.ObjectMeta {
	ns := ""
	// If the defaults have a namespace, that is the first preference.
	if s.Namespace != nil {
		ns = *s.Namespace
	}
	// If the service has a namespace, that is the second preference.
	if service.Namespace != nil {
		ns = *service.Namespace
	}
	// If neither, use the service name.
	if ns == "" {
		ns = *service.Name
	}

	// Append stage to the namespace.
	ns = ns + "-" + string(s.Stage)

	return &metav1.ObjectMeta{
		Name:      *service.Name,
		Namespace: ns,
	}
}
