load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "jsonparser",
    srcs = [
        "bytes.go",
        "bytes_unsafe.go",
        "escape.go",
        "fuzz.go",
        "parser.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/buger/jsonparser",
    importpath = "github.com/buger/jsonparser",
    visibility = ["//visibility:public"],
)
