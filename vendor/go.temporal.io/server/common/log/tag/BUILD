load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tag",
    srcs = [
        "interface.go",
        "tags.go",
        "values.go",
        "zap_tag.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/log/tag",
    importpath = "go.temporal.io/server/common/log/tag",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.uber.org/zap",
        "//vendor/go.uber.org/zap/zapcore",
    ],
)
