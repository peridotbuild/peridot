load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cron",
    srcs = [
        "chain.go",
        "constantdelay.go",
        "cron.go",
        "doc.go",
        "logger.go",
        "option.go",
        "parser.go",
        "spec.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/robfig/cron/v3",
    importpath = "github.com/robfig/cron/v3",
    visibility = ["//visibility:public"],
)
