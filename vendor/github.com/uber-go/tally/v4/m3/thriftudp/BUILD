load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "thriftudp",
    srcs = [
        "multitransport.go",
        "transport.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/m3/thriftudp",
    importpath = "github.com/uber-go/tally/v4/m3/thriftudp",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/uber-go/tally/v4/thirdparty/github.com/apache/thrift/lib/go/thrift",
        "//vendor/go.uber.org/atomic",
    ],
)
