load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go",
    srcs = [
        "db.go",
        "flags.go",
        "frontend_server.go",
        "grpc.go",
        "log.go",
        "pointer.go",
        "slice.go",
    ],
    importpath = "go.resf.org/peridot/base/go",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware",
        "//vendor/github.com/grpc-ecosystem/go-grpc-prometheus",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/github.com/wk8/go-ordered-map/v2:go-ordered-map",
        "//vendor/go.ciq.dev/pika",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials/insecure",
    ],
)
