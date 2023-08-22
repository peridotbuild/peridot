load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0123",
    srcs = [
        "aip0123.go",
        "duplicate_resource.go",
        "name_never_optional.go",
        "resource_annotation.go",
        "resource_definition_pattern.go",
        "resource_definition_type_name.go",
        "resource_definition_variables.go",
        "resource_name_components_alternate.go",
        "resource_name_field.go",
        "resource_pattern.go",
        "resource_plural.go",
        "resource_reference_type.go",
        "resource_singular.go",
        "resource_type_name.go",
        "resource_variables.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0123",
    importpath = "github.com/googleapis/api-linter/rules/aip0123",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/stoewer/go-strcase",
        "@go_googleapis//google/api:annotations_go_proto",
        "@io_bazel_rules_go//proto/wkt:descriptor_go_proto",
    ],
)
