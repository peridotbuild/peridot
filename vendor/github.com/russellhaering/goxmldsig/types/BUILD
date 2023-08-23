load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "types",
    srcs = ["signature.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/russellhaering/goxmldsig/types",
    importpath = "github.com/russellhaering/goxmldsig/types",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/beevik/etree"],
)
