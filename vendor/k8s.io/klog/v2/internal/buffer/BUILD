load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "buffer",
    srcs = ["buffer.go"],
    importmap = "go.resf.org/peridot/vendor/k8s.io/klog/v2/internal/buffer",
    importpath = "k8s.io/klog/v2/internal/buffer",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/k8s.io/klog/v2:__subpackages__",
    ],
    deps = ["//vendor/k8s.io/klog/v2/internal/severity"],
)
