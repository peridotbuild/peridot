load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "payload",
    srcs = ["payload.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/payload",
    importpath = "go.temporal.io/server/common/payload",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/sdk/converter",
        "//vendor/go.temporal.io/server/common/util",
        "//vendor/golang.org/x/exp/maps",
    ],
)
