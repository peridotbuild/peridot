load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-hostpool",
    srcs = [
        "epsilon_greedy.go",
        "epsilon_value_calculators.go",
        "host_entry.go",
        "hostpool.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/hailocab/go-hostpool",
    importpath = "github.com/hailocab/go-hostpool",
    visibility = ["//visibility:public"],
)
