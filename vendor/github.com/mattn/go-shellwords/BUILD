load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-shellwords",
    srcs = [
        "shellwords.go",
        "util_posix.go",
        "util_windows.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/mattn/go-shellwords",
    importpath = "github.com/mattn/go-shellwords",
    visibility = ["//visibility:public"],
)
