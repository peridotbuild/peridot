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

def _oapi_gen_ts_sdk_impl(ctx):
    """
    :param ctx: rule context

    Converts an OpenAPI spec to a TypeScript SDK using openapi-generator.
    """
    java_runtime = ctx.attr._jdk[java_common.JavaRuntimeInfo]
    openapi = ctx.file.openapi

    # Run openapi-generator. With typescript we can just export a directory
    # containing the generated code.
    output_dir = ctx.actions.declare_directory(ctx.attr.name)
    index_ts = ctx.actions.declare_file(ctx.attr.name + "/index.ts")
    gen = ctx.actions.declare_file(ctx.attr.name + "_gen.sh")
    ctx.actions.expand_template(
        template = ctx.file._oapi_converter,
        output = gen,
        substitutions = {
            "{{generator_cli_jar}}": ctx.file._oapi_jar.path,
            "{{openapi_file}}": openapi.path,
            "{{generator}}": "typescript-fetch",
            "{{output_dir}}": output_dir.path,
            "{{java_bin}}": java_runtime.java_executable_exec_path,
        },
        is_executable = True,
    )

    ctx.actions.run(
        inputs = [openapi, ctx.file._oapi_jar] + java_runtime.files.to_list(),
        outputs = [output_dir, index_ts],
        executable = gen,
        env = {
            "JAVA_HOME": java_runtime.java_home,
        },
    )

    return [
        DefaultInfo(
            files = depset([output_dir, index_ts]),
        ),
    ]

oapi_gen_ts_sdk = rule(
    implementation = _oapi_gen_ts_sdk_impl,
    attrs = {
        "openapi": attr.label(
            allow_single_file = True,
            mandatory = True,
            doc = "The OpenAPI spec for the API, can be generated with protoc_gen_openapiv2",
        ),
        "_jdk": attr.label(
            default = Label("@bazel_tools//tools/jdk:current_host_java_runtime"),
            providers = [java_common.JavaRuntimeInfo],
        ),
        "_oapi_jar": attr.label(
            allow_single_file = True,
            default = Label("//third_party:openapi-generator-cli-6.6.0.jar"),
            doc = "The openapi-generator CLI jar",
        ),
        "_oapi_converter": attr.label(
            allow_single_file = True,
            default = Label("//tools/build_rules/oapi_gen:oapi.sh"),
            doc = "The script that converts an OpenAPI spec to a TypeScript SDK",
        ),
    },
)
