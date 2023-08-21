load("@aspect_rules_swc//swc:defs.bzl", "swc")
load("@aspect_rules_esbuild//esbuild:defs.bzl", "esbuild")

def ui_bundle(name, srcs = [], data = [], deps = [], css = False):
    swc(
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
        ],
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
