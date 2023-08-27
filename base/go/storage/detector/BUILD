load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "detector",
    srcs = ["detector.go"],
    importpath = "go.resf.org/peridot/base/go/storage/detector",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//base/go/storage",
        "//base/go/storage/s3",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)
