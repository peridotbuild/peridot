load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "checks",
    srcs = [
        "custom.go",
        "dns.go",
        "http.go",
        "must.go",
        "ping.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/AppsFlyer/go-sundheit/checks",
    importpath = "github.com/AppsFlyer/go-sundheit/checks",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/AppsFlyer/go-sundheit",
        "//vendor/github.com/pkg/errors",
    ],
)
