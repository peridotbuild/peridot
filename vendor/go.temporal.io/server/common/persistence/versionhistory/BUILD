load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "versionhistory",
    srcs = [
        "version_histories.go",
        "version_history.go",
        "version_history_item.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/versionhistory",
    importpath = "go.temporal.io/server/common/persistence/versionhistory",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/history/v1:history",
        "//vendor/go.temporal.io/server/common",
    ],
)
