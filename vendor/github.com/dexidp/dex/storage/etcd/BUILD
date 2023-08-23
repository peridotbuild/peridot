load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "etcd",
    srcs = [
        "config.go",
        "etcd.go",
        "types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/etcd",
    importpath = "github.com/dexidp/dex/storage/etcd",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/dexidp/dex/storage",
        "//vendor/go.etcd.io/etcd/client/pkg/v3/transport",
        "//vendor/go.etcd.io/etcd/client/v3:client",
        "//vendor/go.etcd.io/etcd/client/v3/namespace",
        "//vendor/gopkg.in/square/go-jose.v2:go-jose_v2",
    ],
)
