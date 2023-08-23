load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "kubernetes",
    srcs = [
        "client.go",
        "doc.go",
        "lock.go",
        "storage.go",
        "transport.go",
        "types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/kubernetes",
    importpath = "github.com/dexidp/dex/storage/kubernetes",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Masterminds/semver",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/dexidp/dex/storage",
        "//vendor/github.com/dexidp/dex/storage/kubernetes/k8sapi",
        "//vendor/github.com/ghodss/yaml",
        "//vendor/gopkg.in/square/go-jose.v2:go-jose_v2",
        "@org_golang_x_net//http2",
    ],
)
