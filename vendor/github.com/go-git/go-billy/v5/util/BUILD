load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = [
        "glob.go",
        "util.go",
        "walk.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/go-git/go-billy/v5/util",
    importpath = "github.com/go-git/go-billy/v5/util",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/go-git/go-billy/v5:go-billy"],
)
