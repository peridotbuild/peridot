load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "serialize",
    srcs = ["keyvalues.go"],
    importmap = "go.resf.org/peridot/vendor/k8s.io/klog/v2/internal/serialize",
    importpath = "k8s.io/klog/v2/internal/serialize",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/k8s.io/klog/v2:__subpackages__",
    ],
    deps = ["//vendor/github.com/go-logr/logr"],
)
