load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "interceptor",
    srcs = [
        "interceptor.go",
        "tracing_interceptor.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/sdk/interceptor",
    importpath = "go.temporal.io/sdk/interceptor",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/sdk/activity",
        "//vendor/go.temporal.io/sdk/client",
        "//vendor/go.temporal.io/sdk/converter",
        "//vendor/go.temporal.io/sdk/internal",
        "//vendor/go.temporal.io/sdk/log",
        "//vendor/go.temporal.io/sdk/workflow",
    ],
)
