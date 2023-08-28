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
    name = "mship_worker_server_lib",
    srcs = ["main.go"],
    embedsrcs = ["rh_public_key.asc"],
    importpath = "go.resf.org/peridot/tools/mothership/cmd/mship_worker_server",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//base/go/storage/detector",
        "//tools/mothership/worker_server",
        "//tools/mothership/worker_server/forge",
        "//tools/mothership/worker_server/forge/github",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/go.temporal.io/sdk/client",
        "//vendor/go.temporal.io/sdk/worker",
        "//vendor/golang.org/x/crypto/openpgp",
    ],
)

go_binary(
    name = "mship_worker_server",
    embed = [":mship_worker_server_lib"],
    visibility = ["//visibility:public"],
)
