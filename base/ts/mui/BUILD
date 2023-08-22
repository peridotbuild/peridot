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

load("@aspect_rules_swc//swc:defs.bzl", "swc")
load("@aspect_rules_js//js:defs.bzl", "js_library")

swc(
    name = "lib",
    srcs = glob([
        "*.tsx",
        "*.ts",
    ]),
    data = [
        "//:node_modules/@mui/material",
        "//:node_modules/react",
        "//:node_modules/react-dom",
        "//:node_modules/tslib",
    ],
    visibility = ["//visibility:private"],
)

js_library(
    name = "mui",
    srcs = [":lib"],
    visibility = ["//visibility:public"],
)
