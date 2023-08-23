load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "oauth2client",
    srcs = [
        "oauth2client.go",
        "where.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/db/oauth2client",
    importpath = "github.com/dexidp/dex/storage/ent/db/oauth2client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect/sql",
        "//vendor/github.com/dexidp/dex/storage/ent/db/predicate",
    ],
)
