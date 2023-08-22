load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "metrics",
    srcs = ["metrics.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/proto/otlp/metrics/v1",
    importpath = "go.opentelemetry.io/proto/otlp/metrics/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/proto/otlp/common/v1:common",
        "//vendor/go.opentelemetry.io/proto/otlp/resource/v1:resource",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
