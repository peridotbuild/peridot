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
	"go.starlark.net/starlark"
)

// builtinNamespace is a Starlark function that declares a namespace for the
// deployment. Use if it's desired to group all services in a namespace.
func builtinNamespace(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	// Need exactly one argument.
	if args.Len() != 1 {
		return nil, errors.New("namespace expects exactly one argument")
	}

	// Extract namespace from args.
	namespace, ok := args.Index(0).(starlark.String)
	if !ok {
		return nil, errors.New("namespace expects string argument")
	}

	// Set namespace in thread locals.
	thread.SetLocal("namespace", namespace.GoString())

	return starlark.None, nil
}
