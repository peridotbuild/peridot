load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "k8sapi",
    srcs = [
        "client.go",
        "crd_extensions.go",
        "doc.go",
        "extensions.go",
        "time.go",
        "unversioned.go",
        "v1.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/kubernetes/k8sapi",
    importpath = "github.com/dexidp/dex/storage/kubernetes/k8sapi",
    visibility = ["//visibility:public"],
)
