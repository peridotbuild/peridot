load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ringpop-go",
    srcs = [
        "errors.go",
        "handlers.go",
        "options.go",
        "ringpop.go",
        "stats_handler.go",
        "util.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ringpop-go",
    importpath = "github.com/temporalio/ringpop-go",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/apache/thrift/lib/go/thrift",
        "//vendor/github.com/benbjohnson/clock",
        "//vendor/github.com/dgryski/go-farm",
        "//vendor/github.com/temporalio/ringpop-go/events",
        "//vendor/github.com/temporalio/ringpop-go/forward",
        "//vendor/github.com/temporalio/ringpop-go/hashring",
        "//vendor/github.com/temporalio/ringpop-go/logging",
        "//vendor/github.com/temporalio/ringpop-go/membership",
        "//vendor/github.com/temporalio/ringpop-go/shared",
        "//vendor/github.com/temporalio/ringpop-go/swim",
        "//vendor/github.com/temporalio/ringpop-go/util",
        "//vendor/github.com/temporalio/tchannel-go",
        "//vendor/github.com/temporalio/tchannel-go/json",
        "//vendor/github.com/uber-common/bark",
        "@org_golang_x_net//context",
    ],
)
