load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fxevent",
    srcs = [
        "console.go",
        "event.go",
        "logger.go",
        "zap.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.uber.org/fx/fxevent",
    importpath = "go.uber.org/fx/fxevent",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.uber.org/zap",
        "//vendor/go.uber.org/zap/zapcore",
    ],
)
