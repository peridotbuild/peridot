load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "prometheus",
    srcs = [
        "config.go",
        "reporter.go",
        "sanitize.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/prometheus",
    importpath = "github.com/uber-go/tally/v4/prometheus",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/prometheus/client_golang/prometheus",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
        "//vendor/github.com/uber-go/tally/v4:tally",
    ],
)
