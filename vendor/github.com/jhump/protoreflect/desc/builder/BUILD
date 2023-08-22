load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "builder",
    srcs = [
        "builder.go",
        "doc.go",
        "enum.go",
        "field.go",
        "file.go",
        "message.go",
        "resolver.go",
        "service.go",
        "types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/desc/builder",
    importpath = "github.com/jhump/protoreflect/desc/builder",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/desc/internal",
        "//vendor/github.com/jhump/protoreflect/dynamic",
        "@org_golang_google_protobuf//encoding/protowire",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//reflect/protoregistry",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
