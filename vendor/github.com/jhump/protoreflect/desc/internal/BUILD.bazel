load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "proto3_optional.go",
        "registry.go",
        "source_info.go",
        "util.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/desc/internal",
    importpath = "github.com/jhump/protoreflect/desc/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/jhump/protoreflect/desc:__subpackages__",
    ],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//reflect/protoregistry",
        "@org_golang_google_protobuf//types/descriptorpb",
        "@org_golang_google_protobuf//types/dynamicpb",
    ],
)
