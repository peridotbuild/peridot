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
    name = "kubernetes",
    srcs = [
        "deployment.go",
        "env.go",
        "kubernetes.go",
        "metadata.go",
        "service.go",
        "virtual_service.go",
    ],
    importpath = "go.resf.org/peridot/devtools/garrison/internal/pkg/kubernetes",
    visibility = ["//devtools/garrison:__subpackages__"],
    deps = [
        "//base/go",
        "//devtools/garrison/internal/pkg/starlark",
        "//vendor/k8s.io/api/apps/v1:apps",
        "//vendor/k8s.io/api/core/v1:core",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/util/intstr",
        "//vendor/k8s.io/cli-runtime/pkg/printers",
    ],
)
