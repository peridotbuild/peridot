load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "backoff",
    srcs = [
        "retry.go",
        "retrypolicy.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/sdk/internal/common/backoff",
    importpath = "go.temporal.io/sdk/internal/common/backoff",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.temporal.io/sdk:__subpackages__",
    ],
    deps = [
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/sdk/internal/common/retry",
    ],
)
