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
    name = "mship_dev_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/mothership/cmd/mship_dev",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//tools/mothership/admin/rpc",
        "//tools/mothership/admin/ui",
        "//tools/mothership/rpc",
        "//tools/mothership/ui",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "mship_dev",
    embed = [":mship_dev_lib"],
    visibility = ["//visibility:public"],
)
