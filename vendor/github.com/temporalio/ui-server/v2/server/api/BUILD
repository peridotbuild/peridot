load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api",
    srcs = [
        "handler.go",
        "middleware.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/api",
    importpath = "github.com/temporalio/ui-server/v2/server/api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/gateway",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime",
        "//vendor/github.com/labstack/echo/v4:echo",
        "//vendor/github.com/temporalio/ui-server/v2/server/auth",
        "//vendor/github.com/temporalio/ui-server/v2/server/config",
        "//vendor/github.com/temporalio/ui-server/v2/server/generated/api/workflowservice/v1:workflowservice",
        "//vendor/github.com/temporalio/ui-server/v2/server/rpc",
        "//vendor/github.com/temporalio/ui-server/v2/server/version",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_x_net//context",
    ],
)
