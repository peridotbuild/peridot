load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1_10_0",
    srcs = [
        "doc.go",
        "exception.go",
        "http.go",
        "resource.go",
        "schema.go",
        "trace.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/semconv/v1.10.0",
    importpath = "go.opentelemetry.io/otel/semconv/v1.10.0",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel/attribute",
        "//vendor/go.opentelemetry.io/otel/codes",
        "//vendor/go.opentelemetry.io/otel/semconv/internal",
        "//vendor/go.opentelemetry.io/otel/trace",
    ],
)
