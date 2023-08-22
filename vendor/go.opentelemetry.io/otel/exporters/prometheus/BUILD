load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "prometheus",
    srcs = [
        "prometheus.go",
        "sanitize.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/prometheus",
    importpath = "go.opentelemetry.io/otel/exporters/prometheus",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/prometheus/client_golang/prometheus",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
        "//vendor/go.opentelemetry.io/otel",
        "//vendor/go.opentelemetry.io/otel/attribute",
        "//vendor/go.opentelemetry.io/otel/metric",
        "//vendor/go.opentelemetry.io/otel/sdk/instrumentation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/controller/basic",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export/aggregation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/number",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
        "//vendor/go.opentelemetry.io/otel/sdk/resource",
    ],
)
