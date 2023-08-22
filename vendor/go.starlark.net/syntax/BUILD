load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "syntax",
    srcs = [
        "options.go",
        "parse.go",
        "quote.go",
        "scan.go",
        "syntax.go",
        "walk.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.starlark.net/syntax",
    importpath = "go.starlark.net/syntax",
    visibility = ["//visibility:public"],
)
