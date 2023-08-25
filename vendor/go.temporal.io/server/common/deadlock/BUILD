load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "deadlock",
    srcs = [
        "deadlock.go",
        "fx.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/deadlock",
    importpath = "go.temporal.io/server/common/deadlock",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/internal/goro",
        "//vendor/go.uber.org/fx",
        "@org_golang_google_grpc//health",
    ],
)
