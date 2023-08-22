load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "timer",
    srcs = [
        "gate.go",
        "local_gate.go",
        "remote_gate.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/timer",
    importpath = "go.temporal.io/server/common/timer",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.temporal.io/server/common/clock"],
)
