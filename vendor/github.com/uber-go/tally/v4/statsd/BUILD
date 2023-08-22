load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "statsd",
    srcs = ["reporter.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/statsd",
    importpath = "github.com/uber-go/tally/v4/statsd",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/cactus/go-statsd-client/statsd",
        "//vendor/github.com/uber-go/tally/v4:tally",
    ],
)
