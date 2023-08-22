load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "log",
    srcs = [
        "config.go",
        "interface.go",
        "interface_mock.go",
        "lazy_logger.go",
        "noop_logger.go",
        "panic.go",
        "replay_logger.go",
        "sdk_logger.go",
        "throttle_logger.go",
        "with_logger.go",
        "zap_logger.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/log",
    importpath = "go.temporal.io/server/common/log",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/sdk/log",
        "//vendor/go.temporal.io/sdk/workflow",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/quotas",
        "//vendor/go.uber.org/zap",
        "//vendor/go.uber.org/zap/zapcore",
    ],
)
