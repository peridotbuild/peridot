load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "clockwork",
    srcs = [
        "clockwork.go",
        "context.go",
        "ticker.go",
        "timer.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jonboulle/clockwork",
    importpath = "github.com/jonboulle/clockwork",
    visibility = ["//visibility:public"],
)
