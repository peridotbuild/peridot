load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "otlptracegrpc",
    srcs = [
        "client.go",
        "exporter.go",
        "options.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/otel",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/internal/retry",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlptrace",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlptrace/internal/otlpconfig",
        "//vendor/go.opentelemetry.io/proto/otlp/collector/trace/v1:trace",
        "//vendor/go.opentelemetry.io/proto/otlp/trace/v1:trace",
        "@go_googleapis//google/rpc:errdetails_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//credentials",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
    ],
)
