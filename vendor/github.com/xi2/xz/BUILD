load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "xz",
    srcs = [
        "dec_bcj.go",
        "dec_delta.go",
        "dec_lzma2.go",
        "dec_stream.go",
        "dec_util.go",
        "dec_xz.go",
        "doc.go",
        "reader.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/xi2/xz",
    importpath = "github.com/xi2/xz",
    visibility = ["//visibility:public"],
)
