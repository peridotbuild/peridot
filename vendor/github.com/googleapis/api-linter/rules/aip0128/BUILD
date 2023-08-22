load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0128",
    srcs = [
        "aip0128.go",
        "resource_annotations_field.go",
        "resource_reconciling_behavior.go",
        "resource_reconciling_field.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0128",
    importpath = "github.com/googleapis/api-linter/rules/aip0128",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/desc/builder",
    ],
)
