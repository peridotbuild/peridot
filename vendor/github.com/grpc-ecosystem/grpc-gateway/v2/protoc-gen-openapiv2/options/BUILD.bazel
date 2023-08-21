load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "options",
    srcs = [
        "annotations.pb.go",
        "openapiv2.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/descriptorpb",
        "@org_golang_google_protobuf//types/known/structpb",
    ],
)
