package starlark

import (
	"fmt"
	base "go.resf.org/peridot/base/go"
	"go.starlark.net/starlark"
)

type Default struct {
	Istio     *bool
	AwsRegion *string
}

func (m *Default) String() string {
	return fmt.Sprintf("Default{Istio=%v}", m.Istio)
}

func (m *Default) Type() string {
	return "Default"
}

func (m *Default) Freeze() {}

func (m *Default) Truth() starlark.Bool {
	return true
}

func (m *Default) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", m.Type())
}

// validDefaultKwargs maps to key and type of valid kwargs for the Default function.
var validDefaultKwargs = map[string]starlark.Value{
	"istio":      starlark.Bool(false),
	"aws_region": starlark.String(""),
}

// BuiltinDefault is a Starlark function that declares defaults that affect the
// Starlark program and output.
func builtinDefault(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	if _, ok := thread.Local("default").(*Default); ok {
		return nil, fmt.Errorf("default already declared")
	}

	defaults := &Default{}
	kvs := kwargsToKV(kwargs)
	for _, kv := range kvs.KV {
		starlarkType, ok := validDefaultKwargs[kv.Key]
		if !ok {
			return nil, fmt.Errorf("invalid kwarg %s for Default", kv.Key)
		}

		if kv.Value.Type() != starlarkType.Type() {
			return nil, fmt.Errorf("invalid type for kwarg %s, expected %s, got %s", kv.Key, starlarkType.Type(), kv.Value.Type())
		}

		switch kv.Key {
		case "istio":
			defaults.Istio = base.Pointer(valueToBool(kv.Value))
		case "aws_region":
			defaults.AwsRegion = base.Pointer(valueToString(kv.Value))
		default:
			return nil, fmt.Errorf("invalid kwarg %s for Default", kv.Key)
		}
	}

	thread.SetLocal("default", defaults)

	return starlark.None, nil
}
