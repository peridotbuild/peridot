load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "attribute",
    srcs = [
        "doc.go",
        "encoder.go",
        "iterator.go",
        "key.go",
        "kv.go",
        "set.go",
        "type_string.go",
        "value.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/attribute",
    importpath = "go.opentelemetry.io/otel/attribute",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.opentelemetry.io/otel/internal"],
)
