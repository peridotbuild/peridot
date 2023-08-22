load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "forward",
    srcs = [
        "events.go",
        "forwarder.go",
        "request_sender.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ringpop-go/forward",
    importpath = "github.com/temporalio/ringpop-go/forward",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/temporalio/ringpop-go/events",
        "//vendor/github.com/temporalio/ringpop-go/logging",
        "//vendor/github.com/temporalio/ringpop-go/shared",
        "//vendor/github.com/temporalio/ringpop-go/util",
        "//vendor/github.com/temporalio/tchannel-go",
        "//vendor/github.com/temporalio/tchannel-go/raw",
        "//vendor/github.com/temporalio/tchannel-go/thrift",
        "//vendor/github.com/uber-common/bark",
        "@org_golang_x_net//context",
    ],
)
