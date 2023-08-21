package starlark

import (
	"fmt"
	"go.starlark.net/starlark"
)

type Args struct {
	// key, value tuples
	vals *KVs

	frozen bool
}

func (a *Args) String() string {
	return fmt.Sprintf("Args{vals=%v, frozen=%v}", a.vals, a.frozen)
}

func (a *Args) Type() string {
	return "args"
}

func (a *Args) Freeze() {
	a.frozen = true
}

func (a *Args) Truth() starlark.Bool {
	return len(a.vals.KV) != 0
}

func (a *Args) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", a.Type())
}

func (a *Args) Attr(name string) (starlark.Value, error) {
	val, _ := a.vals.Get(name)
	return val, nil
}

func (a *Args) AttrNames() []string {
	names := make([]string, len(a.vals.KV))
	for i, kv := range a.vals.KV {
		names[i] = kv.Key
	}

	return names
}

// BuiltinArgs fetches values from flags and makes them available to the Starlark program.
func builtinArgs(thread *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	a, ok := thread.Local("args").(*Args)
	if a == nil || !ok {
		return nil, fmt.Errorf("args not declared, run flags() first")
	}

	return a, nil
}
