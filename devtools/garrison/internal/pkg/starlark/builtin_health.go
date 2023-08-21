// Copyright 2023 Peridot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package starlark

import (
	"fmt"
	base "go.resf.org/peridot/base/go"
	"go.starlark.net/starlark"
)

type Health struct {
	Http *HealthHTTP
	Grpc *HealthGRPC

	InitialDelaySeconds *int
	PeriodSeconds       *int
	TimeoutSeconds      *int
	SuccessThreshold    *int
	FailureThreshold    *int
}

func (h *Health) String() string {
	return fmt.Sprintf("Health{Http=%v, Grpc=%v}", h.Http, h.Grpc)
}

func (h *Health) Type() string {
	return "Health"
}

func (h *Health) Freeze() {}

func (h *Health) Truth() starlark.Bool {
	return true
}

func (h *Health) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", h.Type())
}

func defaultHealth() *Health {
	return &Health{
		Http: &HealthHTTP{
			Path: "/_ga/healthz",
			Port: 6661,
		},
		InitialDelaySeconds: base.Pointer(1),
		PeriodSeconds:       base.Pointer(3),
		TimeoutSeconds:      base.Pointer(5),
		SuccessThreshold:    base.Pointer(1),
		FailureThreshold:    base.Pointer(10),
	}
}

// validHealthKwargs maps to key and type of valid kwargs for the Health function.
var validHealthKwargs = map[string]starlark.Value{
	"http": &HealthHTTP{},
	"grpc": &HealthGRPC{},

	"initial_delay_seconds": starlark.MakeInt(0),
	"period_seconds":        starlark.MakeInt(0),
	"timeout_seconds":       starlark.MakeInt(0),
	"success_threshold":     starlark.MakeInt(0),
	"failure_threshold":     starlark.MakeInt(0),
}

// builtinHealth is a Starlark function that declares health settings for the target service.
func builtinHealth(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	health := &Health{}

	for _, kwarg := range kwargs {
		key := string(kwarg[0].(starlark.String))
		value := kwarg[1]

		starlarkType, ok := validHealthKwargs[key]
		if !ok {
			return nil, fmt.Errorf("invalid kwarg %s for Health", key)
		}

		if value.Type() != starlarkType.Type() {
			return nil, fmt.Errorf("invalid type for kwarg %s, expected %s, got %s", key, starlarkType.Type(), value.Type())
		}

		switch key {
		case "http":
			health.Http = value.(*HealthHTTP)
			health.Grpc = nil
		case "grpc":
			health.Grpc = value.(*HealthGRPC)
			health.Http = nil
		case "initial_delay_seconds":
			health.InitialDelaySeconds = base.Pointer(valueToInt(value))
		case "period_seconds":
			health.PeriodSeconds = base.Pointer(valueToInt(value))
		case "timeout_seconds":
			health.TimeoutSeconds = base.Pointer(valueToInt(value))
		case "success_threshold":
			health.SuccessThreshold = base.Pointer(valueToInt(value))
		case "failure_threshold":
			health.FailureThreshold = base.Pointer(valueToInt(value))
		}
	}

	// Http or Grpc must be set
	if health.Http == nil && health.Grpc == nil {
		return nil, fmt.Errorf("invalid kwarg http or grpc for Health, expected http or grpc to be set")
	}

	return health, nil
}
