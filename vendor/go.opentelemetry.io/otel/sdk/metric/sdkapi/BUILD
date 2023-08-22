load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sdkapi",
    srcs = [
        "descriptor.go",
        "instrumentkind.go",
        "instrumentkind_string.go",
        "noop.go",
        "sdkapi.go",
        "wrap.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
    importpath = "go.opentelemetry.io/otel/sdk/metric/sdkapi",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/attribute",
        "//vendor/go.opentelemetry.io/otel/metric",
        "//vendor/go.opentelemetry.io/otel/metric/instrument",
        "//vendor/go.opentelemetry.io/otel/metric/instrument/asyncfloat64",
        "//vendor/go.opentelemetry.io/otel/metric/instrument/asyncint64",
        "//vendor/go.opentelemetry.io/otel/metric/instrument/syncfloat64",
        "//vendor/go.opentelemetry.io/otel/metric/instrument/syncint64",
        "//vendor/go.opentelemetry.io/otel/metric/unit",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/number",
    ],
)
