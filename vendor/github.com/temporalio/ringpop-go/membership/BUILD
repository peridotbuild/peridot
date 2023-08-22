load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "membership",
    srcs = [
        "events.go",
        "interface.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ringpop-go/membership",
    importpath = "github.com/temporalio/ringpop-go/membership",
    visibility = ["//visibility:public"],
)
