load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aggregation",
    srcs = [
        "aggregation.go",
        "temporality.go",
        "temporality_string.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/export/aggregation",
    importpath = "go.opentelemetry.io/otel/sdk/metric/export/aggregation",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/sdk/metric/number",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
    ],
)
