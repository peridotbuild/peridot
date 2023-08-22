load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "statsd",
    srcs = ["reporter.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/metrics/tally/statsd",
    importpath = "go.temporal.io/server/common/metrics/tally/statsd",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/cactus/go-statsd-client/statsd",
        "//vendor/github.com/uber-go/tally/v4:tally",
        "//vendor/github.com/uber-go/tally/v4/statsd",
    ],
)
