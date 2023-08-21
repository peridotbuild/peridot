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
