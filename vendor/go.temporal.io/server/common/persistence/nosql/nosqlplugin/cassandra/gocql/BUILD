load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gocql",
    srcs = [
        "batch.go",
        "client.go",
        "consistency.go",
        "errors.go",
        "interfaces.go",
        "interfaces_mock.go",
        "iter.go",
        "query.go",
        "session.go",
        "uuid.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/nosql/nosqlplugin/cassandra/gocql",
    importpath = "go.temporal.io/server/common/persistence/nosql/nosqlplugin/cassandra/gocql",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gocql/gocql",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/auth",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/nosql/nosqlplugin/cassandra/translator",
        "//vendor/go.temporal.io/server/common/resolver",
    ],
)
