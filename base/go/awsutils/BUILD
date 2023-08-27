load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "awsutils",
    srcs = ["config.go"],
    importpath = "go.resf.org/peridot/base/go/awsutils",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws",
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor/github.com/aws/aws-sdk-go/aws/endpoints",
    ],
)
