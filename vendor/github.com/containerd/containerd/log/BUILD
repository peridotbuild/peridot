load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "log",
    srcs = ["context.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/containerd/containerd/log",
    importpath = "github.com/containerd/containerd/log",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/sirupsen/logrus"],
)
