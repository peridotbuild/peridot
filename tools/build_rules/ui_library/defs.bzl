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

def ui_library(name, srcs, visibility = ["//visibility:private"], deps = [], **kwargs):
    swc(
        name = name + "_lib",
        srcs = srcs,
        swcrc = "//:.swcrc",
        visibility = ["//visibility:private"],
        **kwargs
    )

    js_library(
        name = name,
        deps = deps,
        srcs = [":" + name + "_lib"],
        visibility = visibility,
    )
