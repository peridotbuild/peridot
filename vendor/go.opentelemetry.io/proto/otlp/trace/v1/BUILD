load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "trace",
    srcs = ["trace.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/proto/otlp/trace/v1",
    importpath = "go.opentelemetry.io/proto/otlp/trace/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/proto/otlp/common/v1:common",
        "//vendor/go.opentelemetry.io/proto/otlp/resource/v1:resource",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
