load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "modfile",
    srcs = [
        "print.go",
        "read.go",
        "rule.go",
        "work.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/mod/modfile",
    importpath = "golang.org/x/mod/modfile",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/mod/internal/lazyregexp",
        "//vendor/golang.org/x/mod/module",
        "//vendor/golang.org/x/mod/semver",
    ],
)
