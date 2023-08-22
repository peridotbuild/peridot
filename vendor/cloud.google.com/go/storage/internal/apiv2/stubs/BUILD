load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "stubs",
    srcs = ["storage.pb.go"],
    importmap = "go.resf.org/peridot/vendor/cloud.google.com/go/storage/internal/apiv2/stubs",
    importpath = "cloud.google.com/go/storage/internal/apiv2/stubs",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/cloud.google.com/go/storage:__subpackages__",
    ],
    deps = [
        "//vendor/cloud.google.com/go/iam/apiv1/iampb",
        "@go_googleapis//google/api:annotations_go_proto",
        "@go_googleapis//google/type:date_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/emptypb",
        "@org_golang_google_protobuf//types/known/fieldmaskpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)
