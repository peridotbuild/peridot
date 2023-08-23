load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlclient",
    srcs = ["client.go"],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/sqlclient",
    importpath = "ariga.io/atlas/sql/sqlclient",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/ariga.io/atlas/schemahcl",
        "//vendor/ariga.io/atlas/sql/migrate",
        "//vendor/ariga.io/atlas/sql/schema",
    ],
)
