load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "clock",
    srcs = ["clock.go"],
    importmap = "go.resf.org/peridot/vendor/k8s.io/klog/v2/internal/clock",
    importpath = "k8s.io/klog/v2/internal/clock",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/k8s.io/klog/v2:__subpackages__",
    ],
)
