load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0121",
    srcs = [
        "aip0121.go",
        "resource_must_support_get.go",
        "resource_must_support_list.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0121",
    importpath = "github.com/googleapis/api-linter/rules/aip0121",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
