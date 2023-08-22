load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "registry",
    srcs = [
        "doc.go",
        "registry.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/registry",
    importpath = "go.opentelemetry.io/otel/sdk/metric/registry",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/metric/instrument",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
    ],
)
