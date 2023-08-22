load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sdk",
    srcs = [
        "converter.go",
        "factory.go",
        "factory_mock.go",
        "metrics_handler.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/sdk",
    importpath = "go.temporal.io/server/common/sdk",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/sdk/client",
        "//vendor/go.temporal.io/sdk/converter",
        "//vendor/go.temporal.io/sdk/log",
        "//vendor/go.temporal.io/sdk/worker",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/primitives",
    ],
)
