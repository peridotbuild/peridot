load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "otlpmetric",
    srcs = [
        "clients.go",
        "exporter.go",
        "options.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/internal/metrictransform",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export/aggregation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
        "//vendor/go.opentelemetry.io/otel/sdk/resource",
        "//vendor/go.opentelemetry.io/proto/otlp/metrics/v1:metrics",
    ],
)
