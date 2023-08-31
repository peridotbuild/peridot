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

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "rpc",
    srcs = [
        "rescue.go",
        "rpc.go",
        "worker.go",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/admin/rpc",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//tools/mothership/db",
        "//tools/mothership/proto/admin/v1:pb",
        "//tools/mothership/proto/v1:pb",
        "//vendor/go.ciq.dev/pika",
        "//vendor/go.temporal.io/sdk/client",
        "@go_googleapis//google/rpc:errdetails_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//types/known/emptypb",
    ],
)

go_test(
    name = "rpc_test",
    size = "small",
    srcs = [
        "main_test.go",
        "worker_test.go",
    ],
    embed = [":rpc"],
    deps = [
        "//base/go",
        "//tools/mothership/db",
        "//tools/mothership/migrations",
        "//tools/mothership/proto/admin/v1:pb",
        "//vendor/github.com/stretchr/testify/require",
        "//vendor/github.com/testcontainers/testcontainers-go",
        "//vendor/github.com/testcontainers/testcontainers-go/modules/postgres",
        "//vendor/github.com/testcontainers/testcontainers-go/wait",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
    ],
)
