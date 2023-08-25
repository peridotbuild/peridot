load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ringpop",
    srcs = ["ringpop.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/ringpop",
    importpath = "go.temporal.io/server/common/ringpop",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/temporalio/ringpop-go",
        "//vendor/github.com/temporalio/tchannel-go",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/convert",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/membership",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/rpc/encryption",
    ],
)
