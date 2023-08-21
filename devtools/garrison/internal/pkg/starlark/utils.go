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
