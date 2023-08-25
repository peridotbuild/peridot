load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "temporal",
    srcs = [
        "default_failure_converter.go",
        "doc.go",
        "error.go",
        "retry_policy.go",
        "version.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/sdk/temporal",
    importpath = "go.temporal.io/sdk/temporal",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/sdk/converter",
        "//vendor/go.temporal.io/sdk/internal",
    ],
)
