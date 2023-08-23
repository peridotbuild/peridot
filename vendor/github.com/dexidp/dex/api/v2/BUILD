load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api",
    srcs = [
        "api.pb.go",
        "api_grpc.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/api/v2",
    importpath = "github.com/dexidp/dex/api/v2",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
