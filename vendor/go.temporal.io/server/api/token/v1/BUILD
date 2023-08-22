load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "token",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/token/v1",
    importpath = "go.temporal.io/server/api/token/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/go.temporal.io/server/api/clock/v1:clock",
        "//vendor/go.temporal.io/server/api/history/v1:history",
    ],
)
