load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip4232",
    srcs = [
        "aip4232.go",
        "repeated_fields.go",
        "required_fields.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip4232",
    importpath = "github.com/googleapis/api-linter/rules/aip4232",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
