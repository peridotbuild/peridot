load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "opentracing-go",
    srcs = [
        "ext.go",
        "globaltracer.go",
        "gocontext.go",
        "noop.go",
        "propagation.go",
        "span.go",
        "tracer.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/opentracing/opentracing-go",
    importpath = "github.com/opentracing/opentracing-go",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/opentracing/opentracing-go/log"],
)
