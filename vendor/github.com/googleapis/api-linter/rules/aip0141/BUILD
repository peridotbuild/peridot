load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0141",
    srcs = [
        "aip0141.go",
        "count_suffix.go",
        "forbidden_types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0141",
    importpath = "github.com/googleapis/api-linter/rules/aip0141",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/gertd/go-pluralize",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
