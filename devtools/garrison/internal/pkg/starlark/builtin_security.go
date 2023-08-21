package starlark

import (
	"fmt"
	base "go.resf.org/peridot/base/go"
	"go.starlark.net/starlark"
)

type Security struct {
	RunAsUser  *int
	RunAsGroup *int
}

func (s *Security) String() string {
	return fmt.Sprintf("Security{RunAsUser=%v, RunAsGroup=%v}", s.RunAsUser, s.RunAsGroup)
}

func (s *Security) Type() string {
	return "Security"
}

func (s *Security) Freeze() {}

func (s *Security) Truth() starlark.Bool {
	return true
}

func (s *Security) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", s.Type())
}

// validSecurityKwargs maps to key and type of valid kwargs for the Security function.
var validSecurityKwargs = map[string]starlark.Value{
	"run_as_user":  starlark.MakeInt64(0),
	"run_as_group": starlark.MakeInt64(0),
}

// builtinSecurity is a Starlark function that declares security settings for the target service.
func builtinSecurity(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	security := &Security{}

	for _, kwarg := range kwargs {
		key := string(kwarg[0].(starlark.String))
		value := kwarg[1]

		starlarkType, ok := validSecurityKwargs[key]
		if !ok {
			return nil, fmt.Errorf("invalid kwarg %s for Security", key)
		}

		if value.Type() != starlarkType.Type() {
			return nil, fmt.Errorf("invalid type for kwarg %s, expected %s, got %s", key, starlarkType.Type(), value.Type())
		}

		switch key {
		case "run_as_user":
			security.RunAsUser = base.Pointer(valueToInt(value))
		case "run_as_group":
			security.RunAsGroup = base.Pointer(valueToInt(value))
		}
	}

	return security, nil
}
