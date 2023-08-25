load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "envconfig",
    srcs = ["envconfig.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/otlp/internal/envconfig",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/envconfig",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp:__subpackages__",
    ],
    deps = ["//vendor/go.opentelemetry.io/otel/internal/global"],
)
