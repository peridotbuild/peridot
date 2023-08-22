load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "version",
    srcs = [
        "version.go",
        "with_version_header.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/version",
    importpath = "github.com/temporalio/ui-server/v2/server/version",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime",
        "//vendor/github.com/labstack/echo/v4:echo",
        "@org_golang_google_grpc//metadata",
    ],
)
