load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "migrate",
    srcs = [
        "dir.go",
        "lex.go",
        "migrate.go",
    ],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/migrate",
    importpath = "ariga.io/atlas/sql/migrate",
    visibility = ["//visibility:public"],
    deps = ["//vendor/ariga.io/atlas/sql/schema"],
)
