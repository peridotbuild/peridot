load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "configs",
    srcs = ["quotas.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/frontend/configs",
    importpath = "go.temporal.io/server/service/frontend/configs",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/quotas",
    ],
)
