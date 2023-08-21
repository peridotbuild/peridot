package starlark

import (
    "fmt"
    base "go.resf.org/peridot/base/go"
    "go.starlark.net/starlark"
    "strings"
)

type Port struct {
    Name          *string
    ContainerPort *int
    Port          *int
    Expose        *bool
    External      *bool
    Host          *string
    Path          *string
}

func (p *Port) String() string {
    return fmt.Sprintf("Port{ContainerPort=%v, Port=%v}", p.ContainerPort, p.Port)
}

func (p *Port) Type() string {
    return "Port"
}

func (p *Port) Freeze() {}

func (p *Port) Truth() starlark.Bool {
    return true
}

func (p *Port) Hash() (uint32, error) {
    return 0, fmt.Errorf("unhashable type: %s", p.Type())
}

// validPortKwargs maps to key and type of valid kwargs for the Port function.
var validPortKwargs = map[string]starlark.Value{
    "name":           starlark.String(""),
    "container_port": starlark.MakeInt(0),
    "port":           starlark.MakeInt(0),
    "expose":         starlark.Bool(false),
    "external":       starlark.Bool(false),
    "host":           starlark.String(""),
    "path":           starlark.String(""),
    "stage":          starlark.String(""),
}

// builtinPort is a Starlark function that returns a Port. Used to configure exposed ports for a container.
func builtinPort(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
    port := &Port{}

    for _, kwarg := range kwargs {
        key := string(kwarg[0].(starlark.String))
        value := kwarg[1]

        starlarkType, ok := validPortKwargs[key]
        if !ok {
            return nil, fmt.Errorf("invalid kwarg %s for Port", key)
        }

        if value.Type() != starlarkType.Type() {
            return nil, fmt.Errorf("invalid type for kwarg %s, expected %s, got %s", key, starlarkType.Type(), value.Type())
        }

        switch key {
        case "name":
            port.Name = base.Pointer(valueToString(value))
        case "container_port":
            port.ContainerPort = base.Pointer(valueToInt(value))
        case "port":
            port.Port = base.Pointer(valueToInt(value))
        case "expose":
            port.Expose = base.Pointer(valueToBool(value))
        case "external":
            port.External = base.Pointer(valueToBool(value))
        case "host":
            port.Host = base.Pointer(valueToString(value))
        case "path":
            port.Path = base.Pointer(valueToString(value))
        default:
            return nil, fmt.Errorf("invalid kwarg %s for Port", key)
        }
    }

    // Check required kwargs and set defaults
    if port.Name == nil {
        return nil, fmt.Errorf("name kwarg is required for Port")
    }
    if port.Port == nil {
        return nil, fmt.Errorf("port kwarg is required for Port")
    }
    if port.ContainerPort == nil {
        port.ContainerPort = port.Port
    }

    if port.Expose == nil {
        port.Expose = base.Pointer(false)
    }
    if *port.Expose {
        if port.Host == nil {
            return nil, fmt.Errorf("host kwarg is required when expose is true for Port")
        }
        if port.Path == nil {
            port.Path = base.Pointer("/")
        }

        // If not external, then Host must contain ".corp."
        if port.External == nil || !*port.External {
            if !strings.Contains(*port.Host, ".corp.") {
                return nil, fmt.Errorf("host kwarg must contain .corp. when expose is true and external is false for Port")
            }
        } else {
            // If external, then Host must not contain ".corp."
            if strings.Contains(*port.Host, ".corp.") {
                return nil, fmt.Errorf("host kwarg must not contain .corp. when expose is true and external is true for Port")
            }
        }

        // If stage is not prod, then Host must contain ".corp."
        stage := thread.Local("stage").(Stage)
        if stage != StageProduction {
            // Add -{stage} to the first part of the host (as in the part before the first dot)
            parts := strings.Split(*port.Host, ".")
            parts[0] = fmt.Sprintf("%s-%s", parts[0], stage)

            // Add .corp. to the second part of the host (as in the part after the first dot)
            parts[1] = fmt.Sprintf("%s.corp.", parts[1])

            *port.Host = strings.Join(parts, ".")
        }

        // External should also disable BeyondCorp auth checks
        // todo(mustafa): implement this
    }

    return port, nil
}
