load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0148",
    srcs = [
        "aip0148.go",
        "declarative_friendly_fields.go",
        "field_behavior.go",
        "human_names.go",
        "use_uid.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0148",
    importpath = "github.com/googleapis/api-linter/rules/aip0148",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
