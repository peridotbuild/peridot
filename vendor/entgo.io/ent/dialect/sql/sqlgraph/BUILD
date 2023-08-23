load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlgraph",
    srcs = [
        "entql.go",
        "errors.go",
        "graph.go",
    ],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent/dialect/sql/sqlgraph",
    importpath = "entgo.io/ent/dialect/sql/sqlgraph",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect",
        "//vendor/entgo.io/ent/dialect/sql",
        "//vendor/entgo.io/ent/entql",
        "//vendor/entgo.io/ent/schema/field",
    ],
)
