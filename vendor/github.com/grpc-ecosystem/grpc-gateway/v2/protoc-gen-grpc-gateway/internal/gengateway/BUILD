load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gengateway",
    srcs = [
        "doc.go",
        "generator.go",
        "template.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway/internal/gengateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway/internal/gengateway",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/casing",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/generator",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/utilities",
        "@org_golang_google_grpc//grpclog",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/pluginpb",
    ],
)
