load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0131",
    srcs = [
        "aip0131.go",
        "http_body.go",
        "http_method.go",
        "http_uri_name.go",
        "method_signature.go",
        "request_message_name.go",
        "request_name_behavior.go",
        "request_name_field.go",
        "request_name_reference.go",
        "request_name_reference_type.go",
        "request_name_required.go",
        "request_required_fields.go",
        "request_unknown_fields.go",
        "response_message_name.go",
        "synonyms.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0131",
    importpath = "github.com/googleapis/api-linter/rules/aip0131",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
