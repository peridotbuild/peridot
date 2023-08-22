load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0156",
    srcs = [
        "aip0156.go",
        "forbidden_methods.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0156",
    importpath = "github.com/googleapis/api-linter/rules/aip0156",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
