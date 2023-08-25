load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "replicationadmin",
    srcs = [
        "get_dlq.go",
        "merge_dlq.go",
        "purge_dlq.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/api/replicationadmin",
    importpath = "go.temporal.io/server/service/history/api/replicationadmin",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/service/history/consts",
        "//vendor/go.temporal.io/server/service/history/replication",
        "//vendor/go.temporal.io/server/service/history/shard",
    ],
)
