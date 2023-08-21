load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0122",
    srcs = [
        "aip0122.go",
        "camel_case_uris.go",
        "name_suffix.go",
        "resource_collection_identifiers.go",
        "resource_id_output_only.go",
        "resource_reference_type.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0122",
    importpath = "github.com/googleapis/api-linter/rules/aip0122",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/stoewer/go-strcase",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)
