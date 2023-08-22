load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "telemetry",
    srcs = [
        "config.go",
        "grpc.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/telemetry",
    importpath = "go.temporal.io/server/common/telemetry",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
        "//vendor/go.opentelemetry.io/otel/propagation",
        "//vendor/go.opentelemetry.io/otel/sdk/metric/export",
        "//vendor/go.opentelemetry.io/otel/sdk/trace",
        "//vendor/go.opentelemetry.io/otel/trace",
        "//vendor/gopkg.in/yaml.v3:yaml_v3",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//backoff",
        "@org_golang_google_grpc//credentials/insecure",
    ],
)
