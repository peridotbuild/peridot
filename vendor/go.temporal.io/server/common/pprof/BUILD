load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "pprof",
    srcs = [
        "fx.go",
        "pprof.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/pprof",
    importpath = "go.temporal.io/server/common/pprof",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.uber.org/fx",
    ],
)
