load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "murmur3",
    srcs = [
        "murmur.go",
        "murmur128.go",
        "murmur128_amd64.s",
        "murmur128_decl.go",
        "murmur128_gen.go",
        "murmur32.go",
        "murmur32_gen.go",
        "murmur64.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/twmb/murmur3",
    importpath = "github.com/twmb/murmur3",
    visibility = ["//visibility:public"],
)
