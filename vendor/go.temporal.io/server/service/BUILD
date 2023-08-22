load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "service",
    srcs = ["fx.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service",
    importpath = "go.temporal.io/server/service",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/persistence/client",
        "//vendor/go.temporal.io/server/common/rpc",
        "//vendor/go.temporal.io/server/common/rpc/interceptor",
        "//vendor/go.temporal.io/server/common/telemetry",
        "//vendor/go.uber.org/fx",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
