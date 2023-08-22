load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "export",
    srcs = ["metric.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/export",
    importpath = "go.opentelemetry.io/otel/sdk/metric/export",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/attribute",
        "//vendor/go.opentelemetry.io/otel/sdk/instrumentation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/aggregator",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export/aggregation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
        "//vendor/go.opentelemetry.io/otel/sdk/resource",
    ],
)
