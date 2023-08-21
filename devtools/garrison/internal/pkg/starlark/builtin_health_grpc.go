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

type HealthGRPC struct {
	Port int
}

func (h *HealthGRPC) String() string {
	return fmt.Sprintf("HealthGRPC{Port=%v}", h.Port)
}

func (h *HealthGRPC) Type() string {
	return "HealthGRPC"
}

func (h *HealthGRPC) Freeze() {}

func (h *HealthGRPC) Truth() starlark.Bool {
	return true
}

func (h *HealthGRPC) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", h.Type())
}

// validHealthGRPCKwargs maps to key and type of valid kwargs for the HealthGRPC function.
var validHealthGRPCKwargs = map[string]starlark.Value{
	"port": starlark.MakeInt(0),
}

// builtinHealthGRPC is a Starlark function that declares health settings for the target service.
func builtinHealthGRPC(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	http := &HealthGRPC{}

	for _, kwarg := range kwargs {
		key := string(kwarg[0].(starlark.String))
		value := kwarg[1]

		starlarkType, ok := validHealthGRPCKwargs[key]
		if !ok {
			return nil, fmt.Errorf("invalid kwarg %s for HealthGRPC", key)
		}

		if value.Type() != starlarkType.Type() {
			return nil, fmt.Errorf("invalid type for kwarg %s, expected %s, got %s", key, starlarkType.Type(), value.Type())
		}

		switch key {
		case "port":
			http.Port = valueToInt(value)
		}
	}

	// Port must be set
	if http.Port == 0 {
		return nil, fmt.Errorf("invalid kwarg port for HealthGRPC, expected port to be set")
	}

	return http, nil
}
