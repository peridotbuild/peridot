load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "failure",
    srcs = ["failure.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/failure",
    importpath = "go.temporal.io/server/common/failure",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/failure/v1:failure",
    ],
)
