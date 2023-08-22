load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "configs",
    srcs = ["quotas.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/matching/configs",
    importpath = "go.temporal.io/server/service/matching/configs",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.temporal.io/server/common/quotas"],
)
