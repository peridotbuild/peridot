load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sql",
    srcs = [
        "config.go",
        "crud.go",
        "migrate.go",
        "sql.go",
        "sqlite.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/sql",
    importpath = "github.com/dexidp/dex/storage/sql",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/dexidp/dex/storage",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/lib/pq",
        "//vendor/github.com/mattn/go-sqlite3",
    ],
)
