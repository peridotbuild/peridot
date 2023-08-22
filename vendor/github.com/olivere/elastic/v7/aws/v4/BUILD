load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aws",
    srcs = ["aws_v4.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/olivere/elastic/v7/aws/v4",
    importpath = "github.com/olivere/elastic/v7/aws/v4",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor/github.com/aws/aws-sdk-go/aws/signer/v4:signer",
    ],
)
