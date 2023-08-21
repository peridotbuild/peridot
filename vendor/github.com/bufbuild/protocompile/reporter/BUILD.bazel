load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "reporter",
    srcs = [
        "errors.go",
        "reporter.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bufbuild/protocompile/reporter",
    importpath = "github.com/bufbuild/protocompile/reporter",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/bufbuild/protocompile/ast"],
)
