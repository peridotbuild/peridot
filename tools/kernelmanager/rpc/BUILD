load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "rpc",
    srcs = [
        "kernel.go",
        "operation.go",
        "rpc.go",
    ],
    importpath = "go.resf.org/peridot/tools/kernelmanager/rpc",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//base/go/kv",
        "//third_party/googleapis/google/longrunning:longrunning_go_proto",
        "//tools/kernelmanager/proto/v1:pb",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/sdk/client",
        "@go_googleapis//google/rpc:code_go_proto",
        "@go_googleapis//google/rpc:status_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//reflection",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)
