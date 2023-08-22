load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "backoff",
    srcs = [
        "cron.go",
        "jitter.go",
        "retry.go",
        "retrypolicy.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/backoff",
    importpath = "go.temporal.io/server/common/backoff",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/robfig/cron/v3:cron",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/common/convert",
        "//vendor/go.temporal.io/server/common/util",
    ],
)
