load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "rpc",
    srcs = ["rpc.go"],
    importpath = "go.resf.org/peridot/tools/mothership/rpc",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//tools/mothership/proto/v1:pb",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
