load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "http",
    srcs = ["handler.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/AppsFlyer/go-sundheit/http",
    importpath = "github.com/AppsFlyer/go-sundheit/http",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/AppsFlyer/go-sundheit"],
)
