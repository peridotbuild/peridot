load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "otelgrpc",
    srcs = [
        "grpctrace.go",
        "interceptor.go",
        "interceptorinfo.go",
        "semconv.go",
        "version.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
    importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/internal",
        "//vendor/go.opentelemetry.io/otel",
        "//vendor/go.opentelemetry.io/otel/attribute",
        "//vendor/go.opentelemetry.io/otel/baggage",
        "//vendor/go.opentelemetry.io/otel/codes",
        "//vendor/go.opentelemetry.io/otel/propagation",
        "//vendor/go.opentelemetry.io/otel/semconv/v1.12.0:v1_12_0",
        "//vendor/go.opentelemetry.io/otel/trace",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//peer",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//proto",
    ],
)
