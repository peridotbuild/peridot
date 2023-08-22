load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "m3",
    srcs = [
        "config.go",
        "reporter.go",
        "resource_pool.go",
        "sanitize.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/m3",
    importpath = "github.com/uber-go/tally/v4/m3",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/uber-go/tally/v4:tally",
        "//vendor/github.com/uber-go/tally/v4/internal/cache",
        "//vendor/github.com/uber-go/tally/v4/m3/customtransports",
        "//vendor/github.com/uber-go/tally/v4/m3/thrift/v2:thrift",
        "//vendor/github.com/uber-go/tally/v4/m3/thriftudp",
        "//vendor/github.com/uber-go/tally/v4/thirdparty/github.com/apache/thrift/lib/go/thrift",
        "//vendor/go.uber.org/atomic",
    ],
)
