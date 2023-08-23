load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlspec",
    srcs = ["sqlspec.go"],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/sqlspec",
    importpath = "ariga.io/atlas/sql/sqlspec",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/ariga.io/atlas/schemahcl",
        "//vendor/github.com/zclconf/go-cty/cty",
    ],
)
