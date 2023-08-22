load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0163",
    srcs = [
        "aip0163.go",
        "declarative_friendly_required.go",
        "synonyms.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0163",
    importpath = "github.com/googleapis/api-linter/rules/aip0163",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
