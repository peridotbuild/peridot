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
load("//tools/build_rules/ui_bundle:defs.bzl", "ui_bundle")

ui_bundle(
    name = "bundle",
    deps = [
        "//:node_modules/@mui/icons-material",
        "//:node_modules/@mui/material",
        "//base/ts/mui",
        "//tools/mothership/proto/admin/v1:mshipadminpb_ts_proto",
    ],
)

go_library(
    name = "ui",
    srcs = ["ui.go"],
    # keep
    embedsrcs = [
        ":bundle",  # keep
    ],
    importpath = "go.resf.org/peridot/tools/mothership/admin/ui",
    visibility = ["//visibility:public"],
    deps = ["//base/go"],
)
