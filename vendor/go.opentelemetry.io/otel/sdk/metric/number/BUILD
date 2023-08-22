load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "number",
    srcs = [
        "doc.go",
        "kind_string.go",
        "number.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/sdk/metric/number",
    importpath = "go.opentelemetry.io/otel/sdk/metric/number",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.opentelemetry.io/otel/internal"],
)
