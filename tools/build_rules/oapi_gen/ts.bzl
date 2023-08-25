def _oapi_gen_ts_sdk_impl(ctx):
    """
    :param ctx: rule context

    Converts an OpenAPI spec to a TypeScript SDK using openapi-generator.
    """
    openapi = ctx.file.openapi

    # java_binary returns the jar and a wrapper binary that runs it.
    # we need the wrapper binary
    oapi_cli_outputs = ctx.files._oapi_bin
    oapi_cli = ""
    for output in oapi_cli_outputs:
        if output.path.endswith(".jar"):
            continue
        oapi_cli = output.path
        break

    # Run openapi-generator. With typescript we can just export a directory
    # containing the generated code.
    output_dir = ctx.actions.declare_directory(ctx.attr.name)
    index_ts = ctx.actions.declare_file(ctx.attr.name + "/index.ts")
    gen = ctx.actions.declare_file(ctx.attr.name + "_gen.sh")
    ctx.actions.expand_template(
        template = ctx.file._oapi_converter,
        output = gen,
        substitutions = {
            "{{generator_cli_bin}}": oapi_cli,
            "{{openapi_file}}": openapi.path,
            "{{generator}}": "typescript-fetch",
            "{{output_dir}}": output_dir.path,
        },
        is_executable = True,
    )

    ctx.actions.run(
        inputs = [openapi] + ctx.files._oapi_bin,
        outputs = [output_dir, index_ts],
        executable = gen,
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
        "_oapi_bin": attr.label(
            allow_files = True,
            default = Label("//third_party:openapi_generator"),
            doc = "The openapi-generator CLI binary",
        ),
        "_oapi_converter": attr.label(
            allow_single_file = True,
            default = Label("//tools/build_rules/oapi_gen:oapi.sh"),
            doc = "The script that converts an OpenAPI spec to a fern generators.yaml file",
        ),
    },
)
