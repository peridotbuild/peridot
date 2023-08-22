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
    name = "go",
    srcs = [
        "db.go",
        "flags.go",
        "frontend_server.go",
        "grpc.go",
        "log.go",
        "pointer.go",
        "slice.go",
    ],
    importpath = "go.resf.org/peridot/base/go",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware",
        "//vendor/github.com/grpc-ecosystem/go-grpc-prometheus",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/github.com/wk8/go-ordered-map/v2:go-ordered-map",
        "//vendor/go.ciq.dev/pika",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials/insecure",
    ],
)
