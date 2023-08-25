load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "postgres",
    srcs = ["postgres.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/testcontainers/testcontainers-go/modules/postgres",
    importpath = "github.com/testcontainers/testcontainers-go/modules/postgres",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/testcontainers/testcontainers-go"],
)
