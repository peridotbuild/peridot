load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0233",
    srcs = [
        "aip0233.go",
        "http_body.go",
        "http_method.go",
        "http_uri_suffix.go",
        "plural_method_name.go",
        "request_message_name.go",
        "request_parent_field.go",
        "request_parent_reference.go",
        "request_requests_behavior.go",
        "request_requests_field.go",
        "request_unknown_fields.go",
        "resource_reference_type.go",
        "response_message_name.go",
        "response_resource_field.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0233",
    importpath = "github.com/googleapis/api-linter/rules/aip0233",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/gertd/go-pluralize",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/stoewer/go-strcase",
    ],
)
