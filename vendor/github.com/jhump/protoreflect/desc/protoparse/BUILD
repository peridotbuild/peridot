load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "protoparse",
    srcs = [
        "ast.go",
        "doc.go",
        "errors.go",
        "parser.go",
        "resolve_files.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/desc/protoparse",
    importpath = "github.com/jhump/protoreflect/desc/protoparse",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bufbuild/protocompile",
        "//vendor/github.com/bufbuild/protocompile/ast",
        "//vendor/github.com/bufbuild/protocompile/linker",
        "//vendor/github.com/bufbuild/protocompile/options",
        "//vendor/github.com/bufbuild/protocompile/parser",
        "//vendor/github.com/bufbuild/protocompile/protoutil",
        "//vendor/github.com/bufbuild/protocompile/reporter",
        "//vendor/github.com/bufbuild/protocompile/sourceinfo",
        "//vendor/github.com/bufbuild/protocompile/walk",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/desc/protoparse/ast",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
