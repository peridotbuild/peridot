load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "retry",
    srcs = ["retry.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/otlp/internal/retry",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/retry",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp:__subpackages__",
    ],
    deps = ["//vendor/github.com/cenkalti/backoff/v4:backoff"],
)
