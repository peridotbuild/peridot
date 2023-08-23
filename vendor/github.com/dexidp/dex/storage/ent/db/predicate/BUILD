load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "predicate",
    srcs = ["predicate.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/db/predicate",
    importpath = "github.com/dexidp/dex/storage/ent/db/predicate",
    visibility = ["//visibility:public"],
    deps = ["//vendor/entgo.io/ent/dialect/sql"],
)
