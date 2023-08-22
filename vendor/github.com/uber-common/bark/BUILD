load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bark",
    srcs = [
        "cactus_stats_reporter.go",
        "interface.go",
        "logrus_logger.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-common/bark",
    importpath = "github.com/uber-common/bark",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/cactus/go-statsd-client/statsd",
        "//vendor/github.com/sirupsen/logrus",
    ],
)
