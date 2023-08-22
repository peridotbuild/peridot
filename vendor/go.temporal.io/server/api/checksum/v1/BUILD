load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "checksum",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/checksum/v1",
    importpath = "go.temporal.io/server/api/checksum/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/history/v1:history",
    ],
)
