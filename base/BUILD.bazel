load("@bazel_skylib//rules:common_settings.bzl", "bool_flag")

bool_flag(
    name = "minify_esbuild",
    build_setting_default = False,
)

config_setting(
    name = "minify_esbuild_enabled",
    flag_values = {"minify_esbuild": "True"},
)
