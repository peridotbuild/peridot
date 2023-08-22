load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "hashring",
    srcs = [
        "checksummer.go",
        "hashring.go",
        "rbtree.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ringpop-go/hashring",
    importpath = "github.com/temporalio/ringpop-go/hashring",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dgryski/go-farm",
        "//vendor/github.com/temporalio/ringpop-go/events",
        "//vendor/github.com/temporalio/ringpop-go/logging",
        "//vendor/github.com/temporalio/ringpop-go/membership",
        "//vendor/github.com/uber-common/bark",
    ],
)
