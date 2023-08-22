load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "basic",
    srcs = [
        "basic.go",
        "config.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/processor/basic",
    importpath = "go.opentelemetry.io/otel/sdk/metric/processor/basic",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/attribute",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/aggregator",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export/aggregation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
    ],
)
