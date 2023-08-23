load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-sundheit",
    srcs = [
        "check.go",
        "check_listener.go",
        "check_task.go",
        "config.go",
        "health.go",
        "health_listener.go",
        "options.go",
        "types.go",
        "utils.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/AppsFlyer/go-sundheit",
    importpath = "github.com/AppsFlyer/go-sundheit",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/pkg/errors"],
)
