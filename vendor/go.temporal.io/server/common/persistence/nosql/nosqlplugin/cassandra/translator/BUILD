load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "translator",
    srcs = [
        "fixed_address_translator.go",
        "translator_plugin.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/nosql/nosqlplugin/cassandra/translator",
    importpath = "go.temporal.io/server/common/persistence/nosql/nosqlplugin/cassandra/translator",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gocql/gocql",
        "//vendor/go.temporal.io/server/common/config",
    ],
)
