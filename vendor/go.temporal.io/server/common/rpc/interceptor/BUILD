load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "interceptor",
    srcs = [
        "api_name.go",
        "caller_info.go",
        "namespace.go",
        "namespace_count_limit.go",
        "namespace_logger.go",
        "namespace_rate_limit.go",
        "namespace_validator.go",
        "rate_limit.go",
        "retry.go",
        "sdk_version.go",
        "telemetry.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/rpc/interceptor",
    importpath = "go.temporal.io/server/common/rpc/interceptor",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/authorization",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/quotas",
        "//vendor/go.temporal.io/server/common/serviceerror",
        "//vendor/go.temporal.io/version/check",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
