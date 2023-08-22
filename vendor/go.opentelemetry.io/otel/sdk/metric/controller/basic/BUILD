load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "basic",
    srcs = [
        "config.go",
        "controller.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/controller/basic",
    importpath = "go.opentelemetry.io/otel/sdk/metric/controller/basic",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel",
        "//vendor/go.opentelemetry.io/otel/metric",
        "//vendor/go.opentelemetry.io/otel/sdk/instrumentation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/controller/time",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/registry",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/sdkapi",
        "//vendor/go.opentelemetry.io/otel/sdk/resource",
    ],
)
