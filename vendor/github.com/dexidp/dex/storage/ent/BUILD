load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ent",
    srcs = [
        "generate.go",
        "mysql.go",
        "postgres.go",
        "sqlite.go",
        "types.go",
        "utils.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent",
    importpath = "github.com/dexidp/dex/storage/ent",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect/sql",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/dexidp/dex/storage",
        "//vendor/github.com/dexidp/dex/storage/ent/client",
        "//vendor/github.com/dexidp/dex/storage/ent/db",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/lib/pq",
        "//vendor/github.com/mattn/go-sqlite3",
    ],
)
