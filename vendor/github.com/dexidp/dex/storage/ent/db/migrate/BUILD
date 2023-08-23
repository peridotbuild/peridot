load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "migrate",
    srcs = [
        "migrate.go",
        "schema.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/db/migrate",
    importpath = "github.com/dexidp/dex/storage/ent/db/migrate",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect",
        "//vendor/entgo.io/ent/dialect/sql/schema",
        "//vendor/entgo.io/ent/schema/field",
    ],
)
