load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "vclock",
    srcs = ["vclock.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/vclock",
    importpath = "go.temporal.io/server/service/history/vclock",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/clock/v1:clock",
    ],
)
