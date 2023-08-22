load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cache",
    srcs = [
        "cache.go",
        "lru.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/sdk/internal/common/cache",
    importpath = "go.temporal.io/sdk/internal/common/cache",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.temporal.io/sdk:__subpackages__",
    ],
)
