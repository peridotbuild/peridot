load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aggregator",
    srcs = ["aggregator.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/aggregator",
    importpath = "go.opentelemetry.io/otel/sdk/metric/aggregator",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export/aggregation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/number",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
    ],
)
