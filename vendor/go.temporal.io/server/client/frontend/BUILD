load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "frontend",
    srcs = [
        "client.go",
        "client_gen.go",
        "metric_client.go",
        "metric_client_gen.go",
        "retryable_client.go",
        "retryable_client_gen.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/client/frontend",
    importpath = "go.temporal.io/server/client/frontend",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
