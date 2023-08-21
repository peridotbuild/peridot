load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0154",
    srcs = [
        "aip0154.go",
        "declarative_friendly_required.go",
        "field_type.go",
        "no_duplicate_etag.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0154",
    importpath = "github.com/googleapis/api-linter/rules/aip0154",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
