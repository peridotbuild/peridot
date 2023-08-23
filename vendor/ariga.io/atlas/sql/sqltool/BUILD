load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqltool",
    srcs = [
        "doc.go",
        "hidden.go",
        "hidden_windows.go",
        "tool.go",
    ],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/sqltool",
    importpath = "ariga.io/atlas/sql/sqltool",
    visibility = ["//visibility:public"],
    deps = ["//vendor/ariga.io/atlas/sql/migrate"],
)
