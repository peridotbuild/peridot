load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0191",
    srcs = [
        "aip0191.go",
        "csharp_namespace.go",
        "file_layout.go",
        "file_option_consistency.go",
        "filenames.go",
        "java_multiple_files.go",
        "java_outer_classname.go",
        "java_package.go",
        "php_namespace.go",
        "proto_package.go",
        "proto_version.go",
        "ruby_package.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0191",
    importpath = "github.com/googleapis/api-linter/rules/aip0191",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/stoewer/go-strcase",
        "//vendor/golang.org/x/text/cases",
        "//vendor/golang.org/x/text/language",
        "@io_bazel_rules_go//proto/wkt:descriptor_go_proto",
    ],
)
