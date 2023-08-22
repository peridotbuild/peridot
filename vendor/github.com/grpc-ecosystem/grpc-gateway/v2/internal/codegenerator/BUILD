load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "codegenerator",
    srcs = [
        "doc.go",
        "parse_req.go",
        "supported_features.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/codegenerator",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2/internal/codegenerator",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2:__subpackages__",
    ],
    deps = [
        "@org_golang_google_protobuf//compiler/protogen",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/pluginpb",
    ],
)
