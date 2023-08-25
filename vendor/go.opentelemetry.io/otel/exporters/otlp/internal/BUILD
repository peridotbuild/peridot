load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "config.go",
        "header.go",
        "partialsuccess.go",
        "wrappederror.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/otlp/internal",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp:__subpackages__",
    ],
    deps = ["//vendor/go.opentelemetry.io/otel"],
)
