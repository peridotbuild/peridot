load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "common",
    srcs = [
        "interface.go",
        "interface_mock.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/common",
    importpath = "go.temporal.io/server/service/worker/common",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/sdk/worker",
        "//vendor/go.temporal.io/server/common/namespace",
    ],
)
