load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0126",
    srcs = [
        "aip0126.go",
        "unspecified.go",
        "upper_snake_values.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0126",
    importpath = "github.com/googleapis/api-linter/rules/aip0126",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/stoewer/go-strcase",
    ],
)
