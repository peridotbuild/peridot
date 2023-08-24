load("@aspect_rules_swc//swc:defs.bzl", "swc")
load("@aspect_rules_js//js:defs.bzl", "js_library")

def ui_library(name, srcs, visibility = ["//visibility:private"], **kwargs):
    swc(
        name = name + "_lib",
        srcs = srcs,
        swcrc = "//:.swcrc",
        visibility = ["//visibility:private"],
        **kwargs
    )

    js_library(
        name = name,
        srcs = [":" + name + "_lib"],
        visibility = visibility,
    )