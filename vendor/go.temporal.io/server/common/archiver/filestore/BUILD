load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "filestore",
    srcs = [
        "historyArchiver.go",
        "queryParser.go",
        "queryParser_mock.go",
        "util.go",
        "visibilityArchiver.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/archiver/filestore",
    importpath = "go.temporal.io/server/common/archiver/filestore",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dgryski/go-farm",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/xwb1989/sqlparser",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/history/v1:history",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/workflow/v1:workflow",
        "//vendor/go.temporal.io/server/api/archiver/v1:archiver",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/archiver",
        "//vendor/go.temporal.io/server/common/codec",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/convert",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/common/searchattribute",
        "//vendor/go.temporal.io/server/common/util",
    ],
)
