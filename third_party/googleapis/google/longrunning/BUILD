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

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "longrunning_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
        "//:go_gen_grpc_gateway",
    ],
    importpath = "google.golang.org/genproto/googleapis/longrunning",
    proto = "@go_googleapis//google/longrunning:longrunning_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@go_googleapis//google/api:annotations_go_proto",
        "@go_googleapis//google/rpc:status_go_proto",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/known/anypb",
    ],
)

go_library(
    name = "longrunning",
    embed = [
        ":longrunning_go_proto",
    ],
    importpath = "go.resf.org/peridot/third_party/googleapis/google/longrunning",
    visibility = ["//visibility:public"],
)
