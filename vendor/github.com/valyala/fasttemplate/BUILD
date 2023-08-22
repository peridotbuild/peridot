load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fasttemplate",
    srcs = [
        "template.go",
        "unsafe.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/valyala/fasttemplate",
    importpath = "github.com/valyala/fasttemplate",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/valyala/bytebufferpool"],
)
