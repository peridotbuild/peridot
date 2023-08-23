load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "keys",
    srcs = [
        "keys.go",
        "where.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/db/keys",
    importpath = "github.com/dexidp/dex/storage/ent/db/keys",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect/sql",
        "//vendor/github.com/dexidp/dex/storage/ent/db/predicate",
    ],
)
