load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ent",
    srcs = [
        "ent.go",
        "op_string.go",
    ],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent",
    importpath = "entgo.io/ent",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent/schema",
        "//vendor/entgo.io/ent/schema/edge",
        "//vendor/entgo.io/ent/schema/field",
        "//vendor/entgo.io/ent/schema/index",
    ],
)
