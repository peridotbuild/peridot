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

package starlark

import (
	"fmt"
	"go.starlark.net/starlark"
)

type HealthHTTP struct {
	Path string
	Port int
}

func (h *HealthHTTP) String() string {
	return fmt.Sprintf("HealthHTTP{Path=%v}", h.Path)
}

func (h *HealthHTTP) Type() string {
	return "HealthHTTP"
}

func (h *HealthHTTP) Freeze() {}

func (h *HealthHTTP) Truth() starlark.Bool {
	return true
}

func (h *HealthHTTP) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", h.Type())
}

// validHealthHTTPKwargs maps to key and type of valid kwargs for the HealthHTTP function.
var validHealthHTTPKwargs = map[string]starlark.Value{
	"path": starlark.String(""),
	"port": starlark.MakeInt(0),
}

// builtinHealthHTTP is a Starlark function that declares health settings for the target service.
func builtinHealthHTTP(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	http := &HealthHTTP{}

	for _, kwarg := range kwargs {
		key := string(kwarg[0].(starlark.String))
		value := kwarg[1]

		starlarkType, ok := validHealthHTTPKwargs[key]
		if !ok {
			return nil, fmt.Errorf("invalid kwarg %s for HealthHTTP", key)
		}

		if value.Type() != starlarkType.Type() {
			return nil, fmt.Errorf("invalid type for kwarg %s, expected %s, got %s", key, starlarkType.Type(), value.Type())
		}

		switch key {
		case "path":
			http.Path = valueToString(value)
		case "port":
			http.Port = valueToInt(value)
		}
	}

	if http.Path == "" {
		http.Path = "/_ga/healthz"
	}

	if http.Port == 0 {
		http.Port = 6661
	}

	return http, nil
}
