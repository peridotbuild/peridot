load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "openapiconfig",
    srcs = ["openapiconfig.pb.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/openapiconfig",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/openapiconfig",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
