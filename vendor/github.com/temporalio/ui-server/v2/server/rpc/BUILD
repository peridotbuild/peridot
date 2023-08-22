load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "rpc",
    srcs = [
        "rpc.go",
        "tls.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/rpc",
    importpath = "github.com/temporalio/ui-server/v2/server/rpc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/temporalio/ui-server/v2/server/config",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//backoff",
        "@org_golang_google_grpc//credentials",
    ],
)
