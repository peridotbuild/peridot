load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "io",
    srcs = ["ctxio.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/jbenet/go-context/io",
    importpath = "github.com/jbenet/go-context/io",
    visibility = ["//visibility:public"],
    deps = ["@org_golang_x_net//context"],
)
