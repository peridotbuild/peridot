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
