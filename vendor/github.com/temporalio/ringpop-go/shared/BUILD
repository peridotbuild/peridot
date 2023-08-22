load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "shared",
    srcs = [
        "interfaces.go",
        "shared.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ringpop-go/shared",
    importpath = "github.com/temporalio/ringpop-go/shared",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/temporalio/tchannel-go",
        "@org_golang_x_net//context",
    ],
)
