load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "log",
    srcs = [
        "color.go",
        "log.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/labstack/gommon/log",
    importpath = "github.com/labstack/gommon/log",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/labstack/gommon/color",
        "//vendor/github.com/mattn/go-colorable",
        "//vendor/github.com/mattn/go-isatty",
        "//vendor/github.com/valyala/fasttemplate",
    ],
)
