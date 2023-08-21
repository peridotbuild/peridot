load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sourceinfo",
    srcs = [
        "locations.go",
        "registry.go",
        "wrappers.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/desc/sourceinfo",
    importpath = "github.com/jhump/protoreflect/desc/sourceinfo",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/jhump/protoreflect/desc/internal",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//reflect/protodesc",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//reflect/protoregistry",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
