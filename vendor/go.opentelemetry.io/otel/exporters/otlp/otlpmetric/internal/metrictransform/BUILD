load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "metrictransform",
    srcs = [
        "attribute.go",
        "metric.go",
        "resource.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/internal/metrictransform",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/internal/metrictransform",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric:__subpackages__",
    ],
    deps = [
        "//vendor/go.opentelemetry.io/otel/attribute",
        "//vendor/go.opentelemetry.io/otel/sdk/instrumentation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export/aggregation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/number",
        "//vendor/go.opentelemetry.io/otel/sdk/resource",
        "//vendor/go.opentelemetry.io/proto/otlp/common/v1:common",
        "//vendor/go.opentelemetry.io/proto/otlp/metrics/v1:metrics",
        "//vendor/go.opentelemetry.io/proto/otlp/resource/v1:resource",
    ],
)
