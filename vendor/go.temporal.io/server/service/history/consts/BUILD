load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "consts",
    srcs = ["const.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/consts",
    importpath = "go.temporal.io/server/service/history/consts",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/common",
    ],
)
