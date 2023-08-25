load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = ["util.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/util",
    importpath = "go.temporal.io/server/common/util",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/exp/constraints",
        "//vendor/golang.org/x/exp/maps",
    ],
)
