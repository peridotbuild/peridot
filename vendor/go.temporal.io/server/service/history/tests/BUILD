load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tests",
    srcs = ["vars.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/tests",
    importpath = "go.temporal.io/server/service/history/tests",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/namespace/v1:namespace",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/service/history/configs",
    ],
)
