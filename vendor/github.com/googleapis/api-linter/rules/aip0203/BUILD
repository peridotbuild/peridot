load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0203",
    srcs = [
        "aip0203.go",
        "field_behavior_required.go",
        "unordered_list_repeated.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0203",
    importpath = "github.com/googleapis/api-linter/rules/aip0203",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
