load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqljson",
    srcs = [
        "dialect.go",
        "sqljson.go",
    ],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent/dialect/sql/sqljson",
    importpath = "entgo.io/ent/dialect/sql/sqljson",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect",
        "//vendor/entgo.io/ent/dialect/sql",
    ],
)
