load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tablewriter",
    srcs = [
        "csv.go",
        "table.go",
        "table_with_color.go",
        "util.go",
        "wrap.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/olekukonko/tablewriter",
    importpath = "github.com/olekukonko/tablewriter",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/mattn/go-runewidth"],
)
