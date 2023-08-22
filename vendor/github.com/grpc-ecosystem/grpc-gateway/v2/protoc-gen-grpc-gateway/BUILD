load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "protoc-gen-grpc-gateway_lib",
    srcs = ["main.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/codegenerator",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway/internal/gengateway",
        "@org_golang_google_grpc//grpclog",
        "@org_golang_google_protobuf//compiler/protogen",
    ],
)

go_binary(
    name = "protoc-gen-grpc-gateway",
    embed = [":protoc-gen-grpc-gateway_lib"],
    visibility = ["//visibility:public"],
)
