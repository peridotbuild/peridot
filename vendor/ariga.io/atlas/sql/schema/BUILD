load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "schema",
    srcs = [
        "changekind_string.go",
        "dsl.go",
        "inspect.go",
        "migrate.go",
        "schema.go",
    ],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/schema",
    importpath = "ariga.io/atlas/sql/schema",
    visibility = ["//visibility:public"],
)
