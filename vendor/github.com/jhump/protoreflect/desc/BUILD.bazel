load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "desc",
    srcs = [
        "cache.go",
        "convert.go",
        "descriptor.go",
        "descriptor_unsafe.go",
        "doc.go",
        "imports.go",
        "load.go",
        "wrap.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/desc",
    importpath = "github.com/jhump/protoreflect/desc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bufbuild/protocompile/protoutil",
        "//vendor/github.com/jhump/protoreflect/desc/internal",
        "//vendor/github.com/jhump/protoreflect/desc/sourceinfo",
        "//vendor/github.com/jhump/protoreflect/internal",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//reflect/protodesc",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//reflect/protoregistry",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
