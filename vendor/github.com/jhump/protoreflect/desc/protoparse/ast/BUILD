load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ast",
    srcs = [
        "doc.go",
        "enum.go",
        "field.go",
        "file.go",
        "identifiers.go",
        "message.go",
        "no_source.go",
        "node.go",
        "options.go",
        "print.go",
        "ranges.go",
        "service.go",
        "source_pos.go",
        "values.go",
        "walk.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/desc/protoparse/ast",
    importpath = "github.com/jhump/protoreflect/desc/protoparse/ast",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/bufbuild/protocompile/ast"],
)
