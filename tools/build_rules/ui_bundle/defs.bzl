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

load("@aspect_rules_esbuild//esbuild:defs.bzl", "esbuild")
load("//tools/build_rules/ui_library:defs.bzl", "ui_library")

def ui_bundle(name, srcs = [], data = [], deps = [], css = False):
    ui_library(
        name = name + "_lib",
        srcs = native.glob([
            "*.tsx",
            "*/*.tsx",
            "**/*.tsx",
            "*.ts",
            "*/*.ts",
            "**/*.ts",
        ]) + srcs,
    )

    esbuild_kwargs = {}
    if css:
        esbuild_kwargs["output_css"] = name + ".css"
    esbuild(
        name = name,
        srcs = [":" + name + "_lib"],
        data = data,
        deps = deps + [
            "//:node_modules/react",
            "//:node_modules/react-dom",
            "//:node_modules/react-router",
            "//:node_modules/react-router-dom",
            "//:node_modules/tslib",
        ],
        output_dir = True,
        splitting = True,
        entry_point = "entrypoint.tsx",
        # The transpiling is done with swc, so we don't need to do it here.
        target = "esnext",
        sourcemap = "inline",
        minify = select({
            "//base:minify_esbuild_enabled": True,
            "//conditions:default": False,
        }),
        tsconfig = "//:tsconfig",
        **esbuild_kwargs
    )
