load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0216",
    srcs = [
        "aip0216.go",
        "nesting.go",
        "state_field_output_only.go",
        "synonyms.go",
        "value_synonyms.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0216",
    importpath = "github.com/googleapis/api-linter/rules/aip0216",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)
