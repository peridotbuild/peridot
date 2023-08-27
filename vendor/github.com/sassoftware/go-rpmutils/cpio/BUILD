load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cpio",
    srcs = [
        "cpio.go",
        "extract.go",
        "header.go",
        "reader.go",
        "stream.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/sassoftware/go-rpmutils/cpio",
    importpath = "github.com/sassoftware/go-rpmutils/cpio",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/sassoftware/go-rpmutils/fileutil"],
)
