load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "endpoint",
    srcs = ["endpoint.go"],
    importmap = "go.resf.org/peridot/vendor/go.etcd.io/etcd/client/v3/internal/endpoint",
    importpath = "go.etcd.io/etcd/client/v3/internal/endpoint",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.etcd.io/etcd/client/v3:__subpackages__",
    ],
)
