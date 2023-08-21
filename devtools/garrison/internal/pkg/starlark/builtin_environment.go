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
	"go.starlark.net/starlark"
)

// builtinEnvironment is a Starlark function to declare environment variables for a deployment
func builtinEnvironment(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var newKwargs []starlark.Tuple
	// Make sure all kwargs are strings
	for _, kv := range kwargs {
		if _, ok := kv[0].(starlark.String); !ok {
			return nil, &starlark.EvalError{Msg: "environment variables must be strings"}
		}
		if _, ok := kv[1].(starlark.String); !ok {
			return nil, &starlark.EvalError{Msg: "environment variables must be strings"}
		}
		newKwargs = append(newKwargs, kv)
	}

	// Set the dict values
	dict := starlark.NewDict(len(newKwargs))
	for _, kv := range newKwargs {
		err := dict.SetKey(kv[0], kv[1])
		if err != nil {
			return nil, err
		}
	}

	return dict, nil
}
