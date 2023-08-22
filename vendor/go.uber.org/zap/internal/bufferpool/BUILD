load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bufferpool",
    srcs = ["bufferpool.go"],
    importmap = "go.resf.org/peridot/vendor/go.uber.org/zap/internal/bufferpool",
    importpath = "go.uber.org/zap/internal/bufferpool",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.uber.org/zap:__subpackages__",
    ],
    deps = ["//vendor/go.uber.org/zap/buffer"],
)
