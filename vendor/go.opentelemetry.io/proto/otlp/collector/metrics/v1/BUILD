load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "metrics",
    srcs = [
        "metrics_service.pb.go",
        "metrics_service.pb.gw.go",
        "metrics_service_grpc.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/proto/otlp/collector/metrics/v1",
    importpath = "go.opentelemetry.io/proto/otlp/collector/metrics/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/utilities",
        "//vendor/go.opentelemetry.io/proto/otlp/metrics/v1:metrics",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//grpclog",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
