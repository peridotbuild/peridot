load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sum",
    srcs = ["sum.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/aggregator/sum",
    importpath = "go.opentelemetry.io/otel/sdk/metric/aggregator/sum",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/sdk/metric/aggregator",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export/aggregation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/number",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
    ],
)
