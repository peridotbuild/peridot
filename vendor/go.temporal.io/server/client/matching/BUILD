load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "matching",
    srcs = [
        "client.go",
        "client_gen.go",
        "loadbalancer.go",
        "metric_client.go",
        "metric_client_gen.go",
        "retryable_client.go",
        "retryable_client_gen.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/client/matching",
    importpath = "go.temporal.io/server/client/matching",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/taskqueue/v1:taskqueue",
        "//vendor/go.temporal.io/server/api/matchingservice/v1:matchingservice",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/serviceerror",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
