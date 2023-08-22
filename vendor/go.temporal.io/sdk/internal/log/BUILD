load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "log",
    srcs = [
        "default_logger.go",
        "memory_logger.go",
        "noop_logger.go",
        "replay_logger.go",
        "test_reporter.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/sdk/internal/log",
    importpath = "go.temporal.io/sdk/internal/log",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.temporal.io/sdk:__subpackages__",
    ],
    deps = ["//vendor/go.temporal.io/sdk/log"],
)
