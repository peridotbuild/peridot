load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0143",
    srcs = [
        "aip0143.go",
        "standardized_codes.go",
        "string_type.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0143",
    importpath = "github.com/googleapis/api-linter/rules/aip0143",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
