load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0192",
    srcs = [
        "absolute_links.go",
        "aip0192.go",
        "deprecated_comment.go",
        "has_comments.go",
        "no_html.go",
        "no_markdown_headings.go",
        "no_markdown_tables.go",
        "only_leading_comments.go",
        "trademarked_names.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0192",
    importpath = "github.com/googleapis/api-linter/rules/aip0192",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
