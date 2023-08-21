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
	base "go.resf.org/peridot/base/go"
	"go.starlark.net/starlark"
)

type Security struct {
	RunAsUser  *int
	RunAsGroup *int
}

func (s *Security) String() string {
	return fmt.Sprintf("Security{RunAsUser=%v, RunAsGroup=%v}", s.RunAsUser, s.RunAsGroup)
}

func (s *Security) Type() string {
	return "Security"
}

func (s *Security) Freeze() {}

func (s *Security) Truth() starlark.Bool {
	return true
}

func (s *Security) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", s.Type())
}

// validSecurityKwargs maps to key and type of valid kwargs for the Security function.
var validSecurityKwargs = map[string]starlark.Value{
	"run_as_user":  starlark.MakeInt64(0),
	"run_as_group": starlark.MakeInt64(0),
}

// builtinSecurity is a Starlark function that declares security settings for the target service.
func builtinSecurity(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	security := &Security{}

	for _, kwarg := range kwargs {
		key := string(kwarg[0].(starlark.String))
		value := kwarg[1]

		starlarkType, ok := validSecurityKwargs[key]
		if !ok {
			return nil, fmt.Errorf("invalid kwarg %s for Security", key)
		}

		if value.Type() != starlarkType.Type() {
			return nil, fmt.Errorf("invalid type for kwarg %s, expected %s, got %s", key, starlarkType.Type(), value.Type())
		}

		switch key {
		case "run_as_user":
			security.RunAsUser = base.Pointer(valueToInt(value))
		case "run_as_group":
			security.RunAsGroup = base.Pointer(valueToInt(value))
		}
	}

	return security, nil
}
