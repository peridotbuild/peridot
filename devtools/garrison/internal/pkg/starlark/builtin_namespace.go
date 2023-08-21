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
