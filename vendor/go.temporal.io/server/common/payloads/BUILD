load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "payloads",
    srcs = ["payloads.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/payloads",
    importpath = "go.temporal.io/server/common/payloads",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/sdk/converter",
    ],
)
