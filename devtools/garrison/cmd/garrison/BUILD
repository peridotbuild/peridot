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

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "garrison_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/devtools/garrison/cmd/garrison",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//devtools/garrison/internal/pkg/kubernetes",
        "//devtools/garrison/internal/pkg/starlark",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "garrison",
    embed = [":garrison_lib"],
    visibility = ["//visibility:public"],
)
