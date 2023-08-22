load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cron",
    srcs = [
        "constantdelay.go",
        "cron.go",
        "doc.go",
        "parser.go",
        "spec.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/robfig/cron",
    importpath = "github.com/robfig/cron",
    visibility = ["//visibility:public"],
)
