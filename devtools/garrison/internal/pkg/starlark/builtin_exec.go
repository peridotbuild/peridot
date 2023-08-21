package starlark

import (
	"errors"
	"go.starlark.net/starlark"
	"os"
)

// builtinExec allows for the execution of another Starlark file.
// Useful for embedding other Starlark files. (Declaring functions, etc.)
func builtinExec(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	fPath, ok := args.Index(0).(starlark.String)
	if !ok {
		return nil, errors.New("exec expects a string argument")
	}
	filePath := string(fPath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}

	globals, err := starlark.ExecFile(thread, filePath, nil, predeclared)
	if err != nil {
		return nil, err
	}

	// Get the globals local
	globalsLocal, ok := thread.Local("globals").(*starlark.Dict)
	if !ok {
		globalsLocal = starlark.NewDict(len(globals.Keys()))
	}

	// Add the globals to the local globals
	for _, key := range globals.Keys() {
		value := globals[key]
		err := globalsLocal.SetKey(starlark.String(key), value)
		if err != nil {
			return nil, err
		}
	}

	// Set the globals local
	thread.SetLocal("globals", globalsLocal)

	return starlark.None, nil
}
