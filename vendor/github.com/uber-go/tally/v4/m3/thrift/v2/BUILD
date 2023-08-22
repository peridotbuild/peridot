load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "thrift",
    srcs = [
        "constants.go",
        "m3.go",
        "ttypes.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/m3/thrift/v2",
    importpath = "github.com/uber-go/tally/v4/m3/thrift/v2",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/uber-go/tally/v4/thirdparty/github.com/apache/thrift/lib/go/thrift"],
)
