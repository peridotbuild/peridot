load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "baggage",
    srcs = [
        "baggage.go",
        "context.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/internal/baggage",
    importpath = "go.opentelemetry.io/otel/internal/baggage",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel:__subpackages__",
    ],
)
