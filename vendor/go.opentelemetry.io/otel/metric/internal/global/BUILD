load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "global",
    srcs = [
        "instruments.go",
        "meter.go",
        "state.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/metric/internal/global",
    importpath = "go.opentelemetry.io/otel/metric/internal/global",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel/metric:__subpackages__",
    ],
    deps = [
        "//vendor/go.opentelemetry.io/otel",
        "//vendor/go.opentelemetry.io/otel/attribute",
        "//vendor/go.opentelemetry.io/otel/internal/global",
        "//vendor/go.opentelemetry.io/otel/metric",
        "//vendor/go.opentelemetry.io/otel/metric/instrument",
        "//vendor/go.opentelemetry.io/otel/metric/instrument/asyncfloat64",
        "//vendor/go.opentelemetry.io/otel/metric/instrument/asyncint64",
        "//vendor/go.opentelemetry.io/otel/metric/instrument/syncfloat64",
        "//vendor/go.opentelemetry.io/otel/metric/instrument/syncint64",
    ],
)
