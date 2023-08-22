load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "apiconfig",
    srcs = ["apiconfig.pb.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/apiconfig",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/apiconfig",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2:__subpackages__",
    ],
    deps = [
        "@go_googleapis//google/api:annotations_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
