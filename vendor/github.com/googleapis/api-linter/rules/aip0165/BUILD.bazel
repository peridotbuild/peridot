load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0165",
    srcs = [
        "aip0165.go",
        "http_body.go",
        "http_method.go",
        "http_parent_variable.go",
        "http_uri_suffix.go",
        "request_filter_behavior.go",
        "request_filter_field.go",
        "request_force_field.go",
        "request_message_name.go",
        "request_parent_behavior.go",
        "request_parent_field.go",
        "request_parent_reference.go",
        "response_message_name.go",
        "response_purge_count_field.go",
        "response_purge_sample_field.go",
        "response_purge_sample_reference.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0165",
    importpath = "github.com/googleapis/api-linter/rules/aip0165",
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
