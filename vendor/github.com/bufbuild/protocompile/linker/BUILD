load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "linker",
    srcs = [
        "descriptors.go",
        "doc.go",
        "files.go",
        "linker.go",
        "pathkey_unsafe.go",
        "resolve.go",
        "symbols.go",
        "validate.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bufbuild/protocompile/linker",
    importpath = "github.com/bufbuild/protocompile/linker",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bufbuild/protocompile/ast",
        "//vendor/github.com/bufbuild/protocompile/internal",
        "//vendor/github.com/bufbuild/protocompile/parser",
        "//vendor/github.com/bufbuild/protocompile/protoutil",
        "//vendor/github.com/bufbuild/protocompile/reporter",
        "//vendor/github.com/bufbuild/protocompile/walk",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protodesc",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//reflect/protoregistry",
        "@org_golang_google_protobuf//types/descriptorpb",
        "@org_golang_google_protobuf//types/dynamicpb",
    ],
)
