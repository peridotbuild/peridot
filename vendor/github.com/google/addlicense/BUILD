load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "addlicense_lib",
    srcs = [
        "main.go",
        "tmpl.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/google/addlicense",
    importpath = "github.com/google/addlicense",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/bmatcuk/doublestar/v4:doublestar",
        "//vendor/golang.org/x/sync/errgroup",
    ],
)

go_binary(
    name = "addlicense",
    embed = [":addlicense_lib"],
    visibility = ["//visibility:public"],
)
