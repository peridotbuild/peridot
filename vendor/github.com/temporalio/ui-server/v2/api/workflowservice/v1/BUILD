load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "workflowservice",
    srcs = [
        "service.pb.go",
        "service.pb.gw.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/api/workflowservice/v1",
    importpath = "github.com/temporalio/ui-server/v2/api/workflowservice/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/googleapis/google/api",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/utilities",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "@com_github_golang_protobuf//descriptor:go_default_library_gen",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//grpclog",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
    ],
)
