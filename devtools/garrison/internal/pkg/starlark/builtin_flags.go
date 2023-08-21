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
	"flag"
	"fmt"
	"go.starlark.net/starlark"
	"os"
)

// BuiltinFlags is a Starlark function that declares flags. The flags are then parsed
// and made available to the Starlark program.
func builtinFlags(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	localArgs, ok := thread.Local("args").(*Args)

	var a *Args
	if ok {
		a = localArgs
	} else {
		a = &Args{
			vals: NewKVs(),
		}
	}

	flagMap := make(map[string]any)
	kvs := kwargsToKV(kwargs)
	// Parse flags
	for _, kv := range kvs.KV {
		switch v := kv.Value.(type) {
		case starlark.String:
			x := flag.String(kv.Key, v.GoString(), "")
			flagMap[kv.Key] = x
		case starlark.Bool:
			x := flag.Bool(kv.Key, bool(v), "")
			flagMap[kv.Key] = x
		case starlark.Int:
			x := flag.Int(kv.Key, v.Sign(), "")
			flagMap[kv.Key] = x
		case starlark.Float:
			x := flag.Float64(kv.Key, float64(v), "")
			flagMap[kv.Key] = x
		default:
			return nil, fmt.Errorf("unsupported type %T", v)
		}
	}

	// Since this is called from cmd/garrison, we can assume that the index where
	// the "build" word is located + 2 is the start of the flags.
	buildIndex := 2
	for i, arg := range os.Args {
		if arg == "build" {
			buildIndex = i + 2
			break
		}
	}
	flag.CommandLine.Parse(os.Args[buildIndex:])

	// Set values from flagMap
	for _, kv := range kvs.KV {
		switch v := kv.Value.(type) {
		case starlark.String:
			a.vals.Set(kv.Key, starlark.String(*flagMap[kv.Key].(*string)))
		case starlark.Bool:
			a.vals.Set(kv.Key, starlark.Bool(*flagMap[kv.Key].(*bool)))
		case starlark.Int:
			a.vals.Set(kv.Key, starlark.MakeInt(*flagMap[kv.Key].(*int)))
		case starlark.Float:
			a.vals.Set(kv.Key, starlark.Float(*flagMap[kv.Key].(*float64)))
		default:
			return nil, fmt.Errorf("unsupported type %T", v)
		}

		// Check if value is truthy, otherwise error out because flags are "required"
		// if no default
		if x, ok := a.vals.Get(kv.Key); !ok || !bool(x.Truth()) {
			return nil, fmt.Errorf("flag %s is required", kv.Key)
		}
	}

	thread.SetLocal("args", a)

	return starlark.None, nil
}
