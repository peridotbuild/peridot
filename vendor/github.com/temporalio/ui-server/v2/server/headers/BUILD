load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "headers",
    srcs = ["headers.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/headers",
    importpath = "github.com/temporalio/ui-server/v2/server/headers",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime",
        "//vendor/github.com/labstack/echo/v4:echo",
        "//vendor/github.com/temporalio/ui-server/v2/server/api",
        "@org_golang_google_grpc//metadata",
    ],
)
