load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fileutil",
    srcs = [
        "fileutil_unix.go",
        "fileutil_windows.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/sassoftware/go-rpmutils/fileutil",
    importpath = "github.com/sassoftware/go-rpmutils/fileutil",
    visibility = ["//visibility:public"],
)
