load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "json",
    srcs = [
        "decode.go",
        "encode.go",
        "indent.go",
        "scanner.go",
        "stream.go",
        "tags.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/go-jose/go-jose/v3/json",
    importpath = "github.com/go-jose/go-jose/v3/json",
    visibility = ["//visibility:public"],
)
