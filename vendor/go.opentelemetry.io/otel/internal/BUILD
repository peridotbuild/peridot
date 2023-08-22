load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = ["rawhelpers.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/internal",
    importpath = "go.opentelemetry.io/otel/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel:__subpackages__",
    ],
)
