load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0134",
    srcs = [
        "aip0134.go",
        "http_body.go",
        "http_method.go",
        "http_uri_name.go",
        "method_signature.go",
        "request_allow_missing_field.go",
        "request_mask_field.go",
        "request_mask_required.go",
        "request_message_name.go",
        "request_required_fields.go",
        "request_resource_field.go",
        "request_resource_required.go",
        "request_unknown_fields.go",
        "response_lro.go",
        "response_message_name.go",
        "synonyms.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0134",
    importpath = "github.com/googleapis/api-linter/rules/aip0134",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/stoewer/go-strcase",
    ],
)
