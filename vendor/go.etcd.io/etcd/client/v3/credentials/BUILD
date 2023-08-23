load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "credentials",
    srcs = ["credentials.go"],
    importmap = "go.resf.org/peridot/vendor/go.etcd.io/etcd/client/v3/credentials",
    importpath = "go.etcd.io/etcd/client/v3/credentials",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.etcd.io/etcd/api/v3/v3rpc/rpctypes",
        "@org_golang_google_grpc//credentials",
    ],
)
