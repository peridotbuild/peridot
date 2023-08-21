package starlark

import (
	"fmt"
	base "go.resf.org/peridot/base/go"
	"go.starlark.net/starlark"
)

type Service struct {
	Name           *string
	Image          *string
	Ports          []*Port
	Env            map[string]string
	Command        []string
	Args           []string
	Security       *Security
	Health         *Health
	NoAntiAffinity *bool

	Namespace *string
}

func (s *Service) String() string {
	return fmt.Sprintf("Service{Name=%v, Image=%v, Ports=%v, Env=%v, Namespace=%v}", s.Name, s.Image, s.Ports, s.Env, s.Namespace)
}

type Services struct {
	Services []*Service
}

// validServiceKwargs maps to key and type of valid kwargs for the Service function.
var validServiceKwargs = map[string]starlark.Value{
	"name":       starlark.String(""),
	"image":      starlark.String(""),
	"ports":      starlark.NewList([]starlark.Value{}),
	"env":        starlark.NewDict(0),
	"command":    starlark.NewList([]starlark.Value{}),
	"args":       starlark.NewList([]starlark.Value{}),
	"security":   &Security{},
	"health":     &Health{},
	"no_antiaff": starlark.Bool(false),
}

// BuiltinService is a Starlark function that declares metadata about the target service.
func builtinService(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	localServices, ok := thread.Local("services").(*Services)

	var services *Services
	if ok {
		services = localServices
	} else {
		services = &Services{}
	}

	m := &Service{
		Ports: []*Port{},
	}

	kvs := kwargsToKV(kwargs)
	for _, kv := range kvs.KV {
		starlarkType, ok := validServiceKwargs[kv.Key]
		if !ok {
			return nil, fmt.Errorf("invalid kwarg %s for Service", kv.Key)
		}

		if kv.Value.Type() != starlarkType.Type() {
			return nil, fmt.Errorf("invalid type for kwarg %s, expected %s, got %s", kv.Key, starlarkType.Type(), kv.Value.Type())
		}

		switch kv.Key {
		case "name":
			m.Name = base.Pointer(valueToString(kv.Value))
		case "image":
			m.Image = base.Pointer(valueToString(kv.Value))
		case "ports":
			ports, ok := kv.Value.(*starlark.List)
			if !ok {
				return nil, fmt.Errorf("invalid type for kwarg ports, expected List, got %s", kv.Value.Type())
			}

			// Ensure all elements in the list are of type Port.
			var x starlark.Value
			iterator := ports.Iterate()
			for iterator.Next(&x) {
				if _, ok := x.(*Port); !ok {
					return nil, fmt.Errorf("invalid type for kwarg ports, expected Port, got %s", x.Type())
				}
			}

			m.Ports = make([]*Port, ports.Len())
			for i := 0; i < ports.Len(); i++ {
				m.Ports[i] = ports.Index(i).(*Port)
			}
		case "env":
			env, ok := kv.Value.(*starlark.Dict)
			if !ok {
				return nil, fmt.Errorf("invalid type for kwarg env, expected Dict, got %s", kv.Value.Type())
			}

			m.Env = make(map[string]string)
			var x starlark.Value
			for _, key := range env.Keys() {
				x, _, _ = env.Get(key)
				m.Env[key.(starlark.String).GoString()] = x.(starlark.String).GoString()
			}
		case "command":
			command := kv.Value.(*starlark.List)
			m.Command = make([]string, command.Len())
			for i := 0; i < command.Len(); i++ {
				m.Command[i] = command.Index(i).(starlark.String).GoString()
			}
		case "args":
			sArgs := kv.Value.(*starlark.List)
			m.Args = make([]string, sArgs.Len())
			for i := 0; i < sArgs.Len(); i++ {
				m.Args[i] = sArgs.Index(i).(starlark.String).GoString()
			}
		case "security":
			m.Security = kv.Value.(*Security)
		case "health":
			m.Health = kv.Value.(*Health)
		case "no_antiaff":
			m.NoAntiAffinity = base.Pointer(valueToBool(kv.Value))
		default:
			return nil, fmt.Errorf("invalid kwarg %s for Service", kv.Key)
		}
	}

	// Make sure name and image are set
	if m.Name == nil {
		return nil, fmt.Errorf("service name is required")
	}
	if m.Image == nil {
		return nil, fmt.Errorf("service image is required")
	}

	// Set defaults
	if m.Health == nil {
		m.Health = defaultHealth()
	}

	services.Services = append(services.Services, m)
	thread.SetLocal("services", services)

	return starlark.None, nil
}
