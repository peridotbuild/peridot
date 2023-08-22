load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "auth",
    srcs = [
        "tls.go",
        "tlsConfigHelper.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/auth",
    importpath = "go.temporal.io/server/common/auth",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
    ],
)
