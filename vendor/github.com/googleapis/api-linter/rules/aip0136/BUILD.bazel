load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0136",
    srcs = [
        "aip0136.go",
        "declarative_standard_methods_only.go",
        "http_body.go",
        "http_method.go",
        "http_name_variable.go",
        "http_parent_variable.go",
        "http_uri_suffix.go",
        "prepositions.go",
        "verb_noun.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0136",
    importpath = "github.com/googleapis/api-linter/rules/aip0136",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/gertd/go-pluralize",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/data",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/stoewer/go-strcase",
    ],
)
