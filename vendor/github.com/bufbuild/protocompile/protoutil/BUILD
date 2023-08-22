load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "protoutil",
    srcs = ["protos.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bufbuild/protocompile/protoutil",
    importpath = "github.com/bufbuild/protocompile/protoutil",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protodesc",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
