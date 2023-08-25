load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "migrations",
    srcs = ["migrations.go"],
    embedsrcs = ["000001_init.up.sql"],
    importpath = "go.resf.org/peridot/tools/mothership/migrations",
    visibility = ["//visibility:public"],
)
