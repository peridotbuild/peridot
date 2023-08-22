load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "env",
    srcs = ["env.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/internal/env",
    importpath = "go.opentelemetry.io/otel/sdk/internal/env",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel/sdk:__subpackages__",
    ],
    deps = ["//vendor/go.opentelemetry.io/otel/internal/global"],
)
