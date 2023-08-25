load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "protoc-gen-openapiv2_lib",
    srcs = ["main.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/codegenerator",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/internal/genopenapi",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/utilities",
        "@org_golang_google_grpc//grpclog",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/pluginpb",
    ],
)

go_binary(
    name = "protoc-gen-openapiv2",
    embed = [":protoc-gen-openapiv2_lib"],
    visibility = ["//visibility:public"],
)
