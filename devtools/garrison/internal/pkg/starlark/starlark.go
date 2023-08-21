package starlark

import (
	"fmt"
	"go.starlark.net/starlark"
)

type State struct {
	Services *Services
	Default  *Default

	Stage     Stage
	Namespace *string
}

var predeclaredNoExec = starlark.StringDict{
	"var":         starlark.NewBuiltin("var", builtinVar),
	"flags":       starlark.NewBuiltin("flags", builtinFlags),
	"default":     starlark.NewBuiltin("default", builtinDefault),
	"namespace":   starlark.NewBuiltin("namespace", builtinNamespace),
	"service":     starlark.NewBuiltin("service", builtinService),
	"args":        starlark.NewBuiltin("args", builtinArgs),
	"Port":        starlark.NewBuiltin("Port", builtinPort),
	"Environment": starlark.NewBuiltin("Environment", builtinEnvironment),
	"Security":    starlark.NewBuiltin("Security", builtinSecurity),
	"Health":      starlark.NewBuiltin("Health", builtinHealth),
	"HealthHTTP":  starlark.NewBuiltin("HealthHTTP", builtinHealthHTTP),
	"HealthGRPC":  starlark.NewBuiltin("HealthGRPC", builtinHealthGRPC),
}

var predeclared = starlark.StringDict{}

func init() {
	predeclared = predeclaredNoExec
	predeclared["exec"] = starlark.NewBuiltin("exec", builtinExec)
}

func ExecFile(file string, stage Stage) (*State, error) {
	thread := &starlark.Thread{
		Name: "garrison",
	}

	// Set stage in thread locals.
	thread.SetLocal("stage", stage)

	_, err := starlark.ExecFile(thread, file, nil, predeclared)
	if err != nil {
		return nil, err
	}

	// Extract Services and Default from thread locals.
	services, ok := thread.Local("services").(*Services)
	if !ok {
		return nil, fmt.Errorf("services not declared")
	}

	defaults, ok := thread.Local("default").(*Default)
	if !ok {
		defaults = &Default{}
	}

	var ns *string
	namespace, ok := thread.Local("namespace").(string)
	if ok {
		ns = &namespace
	}

	return &State{
		Services:  services,
		Default:   defaults,
		Stage:     stage,
		Namespace: ns,
	}, nil
}
