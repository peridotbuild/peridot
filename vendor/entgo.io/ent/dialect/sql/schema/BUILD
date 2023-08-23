load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "schema",
    srcs = [
        "atlas.go",
        "inspect.go",
        "migrate.go",
        "mysql.go",
        "postgres.go",
        "schema.go",
        "sqlite.go",
        "writer.go",
    ],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent/dialect/sql/schema",
    importpath = "entgo.io/ent/dialect/sql/schema",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/ariga.io/atlas/sql/migrate",
        "//vendor/ariga.io/atlas/sql/mysql",
        "//vendor/ariga.io/atlas/sql/postgres",
        "//vendor/ariga.io/atlas/sql/schema",
        "//vendor/ariga.io/atlas/sql/sqlclient",
        "//vendor/ariga.io/atlas/sql/sqlite",
        "//vendor/ariga.io/atlas/sql/sqltool",
        "//vendor/entgo.io/ent/dialect",
        "//vendor/entgo.io/ent/dialect/entsql",
        "//vendor/entgo.io/ent/dialect/sql",
        "//vendor/entgo.io/ent/schema/field",
    ],
)
