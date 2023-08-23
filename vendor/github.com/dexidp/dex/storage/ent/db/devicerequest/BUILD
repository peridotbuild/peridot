load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "devicerequest",
    srcs = [
        "devicerequest.go",
        "where.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/db/devicerequest",
    importpath = "github.com/dexidp/dex/storage/ent/db/devicerequest",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect/sql",
        "//vendor/github.com/dexidp/dex/storage/ent/db/predicate",
    ],
)
