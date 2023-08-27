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
        "auth.go",
        "db.go",
        "flags.go",
        "frontend_server.go",
        "fs.go",
        "grpc.go",
        "log.go",
        "pb.go",
        "pointer.go",
        "slice.go",
        "temporal.go",
        "wrapper_helpers.go",
    ],
    embedsrcs = ["assets/oh_no_unauthenticated.png"],
    importpath = "go.resf.org/peridot/base/go",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-oidc/v3/oidc",
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware",
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth",
        "//vendor/github.com/grpc-ecosystem/go-grpc-prometheus",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/github.com/wk8/go-ordered-map/v2:go-ordered-map",
        "//vendor/go.ciq.dev/pika",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/sdk/client",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//credentials",
        "@org_golang_google_grpc//credentials/insecure",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//encoding/protojson",
        "@org_golang_google_protobuf//types/known/timestamppb",
        "@org_golang_google_protobuf//types/known/wrapperspb",
        "@org_golang_x_oauth2//:oauth2",
    ],
)
