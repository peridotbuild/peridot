load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "field",
    srcs = [
        "annotation.go",
        "field.go",
        "numeric.go",
        "type.go",
    ],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent/schema/field",
    importpath = "entgo.io/ent/schema/field",
    visibility = ["//visibility:public"],
    deps = ["//vendor/entgo.io/ent/schema"],
)
