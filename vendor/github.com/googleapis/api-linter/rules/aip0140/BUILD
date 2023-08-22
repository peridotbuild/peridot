load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0140",
    srcs = [
        "abbreviations.go",
        "aip0140.go",
        "base64.go",
        "lower_snake.go",
        "numbers.go",
        "prepositions.go",
        "reserved_words.go",
        "underscores.go",
        "uri.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0140",
    importpath = "github.com/googleapis/api-linter/rules/aip0140",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/data",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/desc/builder",
        "//vendor/github.com/stoewer/go-strcase",
        "//vendor/golang.org/x/text/cases",
        "//vendor/golang.org/x/text/language",
    ],
)
