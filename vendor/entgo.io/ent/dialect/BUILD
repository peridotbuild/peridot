load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "dialect",
    srcs = ["dialect.go"],
    importmap = "go.resf.org/peridot/vendor/entgo.io/ent/dialect",
    importpath = "entgo.io/ent/dialect",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/google/uuid"],
)
