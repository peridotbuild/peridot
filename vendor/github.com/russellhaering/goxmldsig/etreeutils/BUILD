load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "etreeutils",
    srcs = [
        "canonicalize.go",
        "namespace.go",
        "sort.go",
        "unmarshal.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/russellhaering/goxmldsig/etreeutils",
    importpath = "github.com/russellhaering/goxmldsig/etreeutils",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/beevik/etree"],
)
