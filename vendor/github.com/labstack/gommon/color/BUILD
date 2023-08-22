load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "color",
    srcs = ["color.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/labstack/gommon/color",
    importpath = "github.com/labstack/gommon/color",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mattn/go-colorable",
        "//vendor/github.com/mattn/go-isatty",
    ],
)
