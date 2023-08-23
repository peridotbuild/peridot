load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "authcode",
    srcs = [
        "authcode.go",
        "where.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/db/authcode",
    importpath = "github.com/dexidp/dex/storage/ent/db/authcode",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect/sql",
        "//vendor/github.com/dexidp/dex/storage/ent/db/predicate",
    ],
)
