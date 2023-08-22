load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0235",
    srcs = [
        "aip0235.go",
        "http_body.go",
        "http_method.go",
        "http_uri_suffix.go",
        "plural_method_name.go",
        "request_message_name.go",
        "request_names_behavior.go",
        "request_names_field.go",
        "request_names_reference.go",
        "request_parent_field.go",
        "request_parent_reference.go",
        "request_requests_behavior.go",
        "request_unknown_fields.go",
        "response_message_name.go",
        "response_resource_field.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0235",
    importpath = "github.com/googleapis/api-linter/rules/aip0235",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/gertd/go-pluralize",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/desc/builder",
    ],
)
