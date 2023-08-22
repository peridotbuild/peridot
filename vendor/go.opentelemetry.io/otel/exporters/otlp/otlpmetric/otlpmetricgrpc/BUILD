load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "otlpmetricgrpc",
    srcs = [
        "client.go",
        "exporter.go",
        "options.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/internal/retry",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/internal/otlpconfig",
        "//vendor/go.opentelemetry.io/proto/otlp/collector/metrics/v1:metrics",
        "//vendor/go.opentelemetry.io/proto/otlp/metrics/v1:metrics",
        "@go_googleapis//google/rpc:errdetails_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//credentials",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
    ],
)
