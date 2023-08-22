load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cassandra",
    srcs = ["visibility_store.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/visibility/store/standard/cassandra",
    importpath = "go.temporal.io/server/common/persistence/visibility/store/standard/cassandra",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/nosql/nosqlplugin/cassandra/gocql",
        "//vendor/go.temporal.io/server/common/persistence/visibility/manager",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store",
        "//vendor/go.temporal.io/server/common/resolver",
    ],
)
