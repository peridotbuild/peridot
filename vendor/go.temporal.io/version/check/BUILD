load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "check",
    srcs = [
        "caller.go",
        "request.go",
        "response.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/version/check",
    importpath = "go.temporal.io/version/check",
    visibility = ["//visibility:public"],
)
