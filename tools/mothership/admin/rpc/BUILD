load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "rpc",
    srcs = [
        "rpc.go",
        "worker.go",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/admin/rpc",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//tools/mothership/db",
        "//tools/mothership/proto/admin/v1:pb",
        "//vendor/go.ciq.dev/pika",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
