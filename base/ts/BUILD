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

load("//tools/build_rules/ui_library:defs.bzl", "ui_library")

ui_library(
    name = "ts",
    srcs = glob([
        "*.tsx",
        "*.ts",
    ]),
    visibility = ["//visibility:public"],
    deps = [
        "//:node_modules/await-to-js",
    ],
)
