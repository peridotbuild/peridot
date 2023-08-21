load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "longrunningpb",
    srcs = ["operations.pb.go"],
    importmap = "go.resf.org/peridot/vendor/cloud.google.com/go/longrunning/autogen/longrunningpb",
    importpath = "cloud.google.com/go/longrunning/autogen/longrunningpb",
    visibility = ["//visibility:public"],
    deps = [
        "@go_googleapis//google/api:annotations_go_proto",
        "@go_googleapis//google/rpc:status_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/descriptorpb",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/emptypb",
    ],
)
