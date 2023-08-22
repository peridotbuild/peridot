load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0164",
    srcs = [
        "aip0164.go",
        "declarative_friendly_required.go",
        "http_body.go",
        "http_method.go",
        "http_uri_suffix.go",
        "request_message_name.go",
        "request_name_behavior.go",
        "request_name_field.go",
        "request_name_reference.go",
        "request_unknown_fields.go",
        "resource_expire_time_field.go",
        "response_lro.go",
        "response_message_name.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0164",
    importpath = "github.com/googleapis/api-linter/rules/aip0164",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
