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
