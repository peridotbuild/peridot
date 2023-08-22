load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cluster",
    srcs = [
        "fx.go",
        "metadata.go",
        "metadata_mock.go",
        "metadata_test_config.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/cluster",
    importpath = "go.temporal.io/server/common/cluster",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/collection",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/internal/goro",
        "//vendor/go.uber.org/fx",
    ],
)
