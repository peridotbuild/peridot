# Copyright 2023 Peridot Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "starlark",
    srcs = [
        "builtin_args.go",
        "builtin_default.go",
        "builtin_environment.go",
        "builtin_exec.go",
        "builtin_flags.go",
        "builtin_health.go",
        "builtin_health_grpc.go",
        "builtin_health_http.go",
        "builtin_namespace.go",
        "builtin_port.go",
        "builtin_security.go",
        "builtin_service.go",
        "builtin_var.go",
        "kv.go",
        "stage.go",
        "starlark.go",
        "utils.go",
    ],
    importpath = "go.resf.org/peridot/devtools/garrison/internal/pkg/starlark",
    visibility = ["//devtools/garrison:__subpackages__"],
    deps = [
        "//base/go",
        "//vendor/go.starlark.net/starlark",
    ],
)
