load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0158",
    srcs = [
        "aip0158.go",
        "request_page_size_field.go",
        "request_page_token_field.go",
        "request_skip_field.go",
        "response_next_page_token_field.go",
        "response_plural_first_field.go",
        "response_repeated_first_field.go",
        "response_unary.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0158",
    importpath = "github.com/googleapis/api-linter/rules/aip0158",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/desc/builder",
        "//vendor/github.com/stoewer/go-strcase",
    ],
)
