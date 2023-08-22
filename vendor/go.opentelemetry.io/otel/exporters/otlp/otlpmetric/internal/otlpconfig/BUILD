load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "otlpconfig",
    srcs = [
        "envconfig.go",
        "options.go",
        "optiontypes.go",
        "tls.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/internal/otlpconfig",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/internal/otlpconfig",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/otlpmetric:__subpackages__",
    ],
    deps = [
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/internal",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/internal/envconfig",
        "//vendor/go.opentelemetry.io/otel/exporters/otlp/internal/retry",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//backoff",
        "@org_golang_google_grpc//credentials",
        "@org_golang_google_grpc//credentials/insecure",
        "@org_golang_google_grpc//encoding/gzip",
    ],
)
