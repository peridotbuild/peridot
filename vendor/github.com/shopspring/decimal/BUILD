load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "decimal",
    srcs = [
        "decimal.go",
        "decimal-go.go",
        "rounding.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/shopspring/decimal",
    importpath = "github.com/shopspring/decimal",
    visibility = ["//visibility:public"],
)
