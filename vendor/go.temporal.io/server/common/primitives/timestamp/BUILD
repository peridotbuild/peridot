load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "timestamp",
    srcs = [
        "duration.go",
        "parseDuration.go",
        "time.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/primitives/timestamp",
    importpath = "go.temporal.io/server/common/primitives/timestamp",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/go.temporal.io/server/common/util",
    ],
)
