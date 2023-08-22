load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "customtransports",
    srcs = [
        "buffered_read_transport.go",
        "m3_calc_transport.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/m3/customtransports",
    importpath = "github.com/uber-go/tally/v4/m3/customtransports",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/uber-go/tally/v4/thirdparty/github.com/apache/thrift/lib/go/thrift"],
)
