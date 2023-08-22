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

load("@aspect_bazel_lib//lib:copy_to_bin.bzl", "copy_to_bin")
load("@io_bazel_rules_go//proto:compiler.bzl", "go_proto_compiler")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "node_modules")

exports_files(["tsconfig.json"])

# gazelle:prefix go.resf.org/peridot
# gazelle:build_file_name BUILD
# gazelle:go_visibility //third_party:__subpackages__
# gazelle:exclude third_party/googleapis
# gazelle:exclude vendor/go.resf.org/peridot
# gazelle:exclude vendor/google.golang.org
# gazelle:exclude vendor/github.com/golang/protobuf
# gazelle:exclude vendor/golang.org/x/net
# gazelle:exclude vendor/golang.org/x/oauth2
# gazelle:exclude vendor/github.com/googleapis/gax-go/v2
# gazelle:exclude vendor.go
# gazelle:go_grpc_compilers @io_bazel_rules_go//proto:go_grpc,//:go_gen_grpc_gateway
# gazelle:resolve go github.com/bazelbuild/bazel-watcher/internal/ibazel/profiler //third_party/github.com/bazelbuild/bazel-watcher/internal/ibazel/profiler
# gazelle:resolve proto proto google/api/annotations.proto @googleapis//google/api:annotations_proto
# gazelle:resolve proto go google/api/annotations.proto  @org_golang_google_genproto//googleapis/api/annotations
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_proto_compiler(
    name = "go_gen_grpc_gateway",
    options = ["logtostderr=true"],
    plugin = "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway",
    suffix = ".pb.gw.go",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/utilities",
        "@org_golang_google_grpc//grpclog",
        "@org_golang_google_grpc//metadata",
    ],
)

copy_to_bin(
    name = "tsconfig",
    srcs = ["tsconfig.json"],
    visibility = ["//visibility:public"],
)
