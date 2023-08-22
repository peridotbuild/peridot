load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "statsd",
    srcs = [
        "buffer_pool.go",
        "client.go",
        "client_config.go",
        "client_legacy.go",
        "client_noop_legacy.go",
        "doc.go",
        "sender.go",
        "sender_buffered.go",
        "sender_resolving.go",
        "validator.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/cactus/go-statsd-client/statsd",
    importpath = "github.com/cactus/go-statsd-client/statsd",
    visibility = ["//visibility:public"],
)
