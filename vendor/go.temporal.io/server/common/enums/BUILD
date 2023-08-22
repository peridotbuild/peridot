load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "enums",
    srcs = ["defaults.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/enums",
    importpath = "go.temporal.io/server/common/enums",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.temporal.io/api/enums/v1:enums"],
)
