load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "resolver",
    srcs = ["resolver.go"],
    importmap = "go.resf.org/peridot/vendor/go.etcd.io/etcd/client/v3/internal/resolver",
    importpath = "go.etcd.io/etcd/client/v3/internal/resolver",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.etcd.io/etcd/client/v3:__subpackages__",
    ],
    deps = [
        "//vendor/go.etcd.io/etcd/client/v3/internal/endpoint",
        "@org_golang_google_grpc//resolver",
        "@org_golang_google_grpc//resolver/manual",
        "@org_golang_google_grpc//serviceconfig",
    ],
)
