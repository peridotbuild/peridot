load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tlsconfig",
    srcs = [
        "certpool_go17.go",
        "certpool_other.go",
        "config.go",
        "config_client_ciphers.go",
        "config_legacy_client_ciphers.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/go-connections/tlsconfig",
    importpath = "github.com/docker/go-connections/tlsconfig",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/pkg/errors"],
)
