load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0127",
    srcs = [
        "aip0127.go",
        "http_annotation.go",
        "http_template_pattern.go",
        "http_template_syntax.go",
        "resource_name_extraction.go",
        "uri_leading_slash.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0127",
    importpath = "github.com/googleapis/api-linter/rules/aip0127",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
