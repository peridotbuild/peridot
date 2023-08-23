load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "entsql",
    srcs = ["annotation.go"],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent/dialect/entsql",
    importpath = "entgo.io/ent/dialect/entsql",
    visibility = ["//visibility:public"],
    deps = ["//vendor/entgo.io/ent/schema"],
)
