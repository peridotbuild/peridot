load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "checksum",
    srcs = ["crc.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/checksum",
    importpath = "go.temporal.io/server/common/checksum",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
    ],
)
