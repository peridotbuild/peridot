load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sql",
    srcs = [
        "builder.go",
        "driver.go",
        "scan.go",
        "sql.go",
    ],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent/dialect/sql",
    importpath = "entgo.io/ent/dialect/sql",
    visibility = ["//visibility:public"],
    deps = ["//vendor/entgo.io/ent/dialect"],
)
