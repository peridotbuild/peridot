load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ext",
    srcs = [
        "field.go",
        "tags.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/opentracing/opentracing-go/ext",
    importpath = "github.com/opentracing/opentracing-go/ext",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/opentracing/opentracing-go",
        "//vendor/github.com/opentracing/opentracing-go/log",
    ],
)
