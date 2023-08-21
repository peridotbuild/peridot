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
	"errors"
	"fmt"
	"go.starlark.net/starlark"
)

// builtinVar is to fetch a global variable that was defined as a result
// of an exec call
func builtinVar(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	key, ok := args.Index(0).(starlark.String)
	if !ok {
		return nil, errors.New("global expects a string argument")
	}

	globalsLocal, ok := thread.Local("globals").(*starlark.Dict)
	if !ok {
		return nil, fmt.Errorf("global %s not found", key)
	}

	value, ok, err := globalsLocal.Get(key)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("global %s not found", key)
	}

	return value, nil
}
