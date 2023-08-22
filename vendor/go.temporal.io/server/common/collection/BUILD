load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "collection",
    srcs = [
        "concurrent_tx_map.go",
        "interface.go",
        "iterator.go",
        "pagingIterator.go",
        "priorityQueue.go",
        "util.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/collection",
    importpath = "go.temporal.io/server/common/collection",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.temporal.io/server/common/util"],
)
