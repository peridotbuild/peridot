load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "parser",
    srcs = [
        "ast.go",
        "clone.go",
        "doc.go",
        "errors.go",
        "lexer.go",
        "parser.go",
        "proto.y.go",
        "result.go",
        "validate.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bufbuild/protocompile/parser",
    importpath = "github.com/bufbuild/protocompile/parser",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bufbuild/protocompile/ast",
        "//vendor/github.com/bufbuild/protocompile/internal",
        "//vendor/github.com/bufbuild/protocompile/reporter",
        "//vendor/github.com/bufbuild/protocompile/walk",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
