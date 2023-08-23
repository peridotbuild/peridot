load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "edge",
    srcs = [
        "annotation.go",
        "edge.go",
    ],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent/schema/edge",
    importpath = "entgo.io/ent/schema/edge",
    visibility = ["//visibility:public"],
    deps = ["//vendor/entgo.io/ent/schema"],
)
