load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "descriptor",
    srcs = [
        "grpc_api_configuration.go",
        "openapi_configuration.go",
        "registry.go",
        "services.go",
        "types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/golang/glog",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/casing",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/codegenerator",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/apiconfig",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/openapiconfig",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/httprule",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options",
        "//vendor/golang.org/x/text/cases",
        "//vendor/golang.org/x/text/language",
        "//vendor/gopkg.in/yaml.v3:yaml_v3",
        "@go_googleapis//google/api:annotations_go_proto",
        "@org_golang_google_protobuf//compiler/protogen",
        "@org_golang_google_protobuf//encoding/protojson",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/descriptorpb",
        "@org_golang_google_protobuf//types/pluginpb",
    ],
)
