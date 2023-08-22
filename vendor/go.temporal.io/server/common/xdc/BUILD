load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "xdc",
    srcs = [
        "nDCHistoryResender.go",
        "nDCHistoryResender_mock.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/xdc",
    importpath = "go.temporal.io/server/common/xdc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/server/api/adminservice/v1:adminservice",
        "//vendor/go.temporal.io/server/api/history/v1:history",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/client",
        "//vendor/go.temporal.io/server/common/collection",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence/serialization",
        "//vendor/go.temporal.io/server/common/rpc",
    ],
)
