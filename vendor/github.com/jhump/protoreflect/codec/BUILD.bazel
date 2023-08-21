load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "codec",
    srcs = [
        "codec.go",
        "decode_fields.go",
        "doc.go",
        "encode_fields.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/codec",
    importpath = "github.com/jhump/protoreflect/codec",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/internal/codec",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
