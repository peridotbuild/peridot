load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0124",
    srcs = [
        "aip0124.go",
        "reference_same_package.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0124",
    importpath = "github.com/googleapis/api-linter/rules/aip0124",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
