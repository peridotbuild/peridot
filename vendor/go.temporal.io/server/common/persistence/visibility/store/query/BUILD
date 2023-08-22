load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "query",
    srcs = [
        "converter.go",
        "errors.go",
        "interceptors.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/visibility/store/query",
    importpath = "go.temporal.io/server/common/persistence/visibility/store/query",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/olivere/elastic/v7:elastic",
        "//vendor/github.com/xwb1989/sqlparser",
        "//vendor/go.temporal.io/api/serviceerror",
    ],
)
