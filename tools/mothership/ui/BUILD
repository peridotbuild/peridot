load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("//tools/build_rules/ui_bundle:defs.bzl", "ui_bundle")

ui_bundle(
    name = "bundle",
    deps = [
        "//:node_modules/@mui/icons-material",
        "//:node_modules/@mui/material",
        "//base/ts/mui",
    ],
)

go_library(
    name = "ui_lib",
    srcs = ["main.go"],
    embedsrcs = [
        ":bundle.js",  # keep
    ],
    importpath = "go.resf.org/peridot/tools/mothership/ui",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "ui",
    embed = [":ui_lib"],
    visibility = ["//visibility:public"],
)
