load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "packager",
    srcs = ["repack.go"],
    importpath = "go.resf.org/peridot/tools/kernelmanager/packager",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/go-git/go-billy/v5:go-billy"],
)
