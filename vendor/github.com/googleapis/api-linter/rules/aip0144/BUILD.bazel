load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0144",
    srcs = [
        "aip0144.go",
        "http_body.go",
        "http_method.go",
        "request_message_name.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0144",
    importpath = "github.com/googleapis/api-linter/rules/aip0144",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
