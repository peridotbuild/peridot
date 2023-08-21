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

import "go.starlark.net/starlark"

func valueToInt(v starlark.Value) int {
	vInt, ok := v.(starlark.Int)
	if !ok {
		return 0
	}
	vInt64, ok := vInt.Int64()
	if !ok {
		return 0
	}

	return int(vInt64)
}

func valueToString(v starlark.Value) string {
	vString, ok := v.(starlark.String)
	if !ok {
		return ""
	}

	return string(vString)
}

func valueToBool(v starlark.Value) bool {
	vBool, ok := v.(starlark.Bool)
	if !ok {
		return false
	}

	return bool(vBool)
}
