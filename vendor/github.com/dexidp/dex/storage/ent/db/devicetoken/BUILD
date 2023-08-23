load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "devicetoken",
    srcs = [
        "devicetoken.go",
        "where.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/db/devicetoken",
    importpath = "github.com/dexidp/dex/storage/ent/db/devicetoken",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/dialect/sql",
        "//vendor/github.com/dexidp/dex/storage/ent/db/predicate",
    ],
)
