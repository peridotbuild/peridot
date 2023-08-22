load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "rpc",
    srcs = [
        "context.go",
        "grpc.go",
        "rpc.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/rpc",
    importpath = "go.temporal.io/server/common/rpc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/status",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/convert",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/rpc/encryption",
        "//vendor/go.temporal.io/server/common/serviceerror",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//backoff",
        "@org_golang_google_grpc//credentials",
        "@org_golang_google_grpc//credentials/insecure",
    ],
)
