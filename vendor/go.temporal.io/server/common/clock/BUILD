load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "clock",
    srcs = ["time_source.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/clock",
    importpath = "go.temporal.io/server/common/clock",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/jonboulle/clockwork"],
)
