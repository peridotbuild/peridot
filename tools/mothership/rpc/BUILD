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
    name = "rpc",
    srcs = [
        "batch.go",
        "entry.go",
        "operation.go",
        "ping.go",
        "rpc.go",
        "worker.go",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/rpc",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//third_party/googleapis/google/longrunning:longrunning_go_proto",
        "//tools/mothership/db",
        "//tools/mothership/proto/v1:pb",
        "//tools/mothership/worker_server",
        "//vendor/go.ciq.dev/pika",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/sdk/client",
        "@go_googleapis//google/rpc:code_go_proto",
        "@go_googleapis//google/rpc:status_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/emptypb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)
