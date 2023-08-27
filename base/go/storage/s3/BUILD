load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "s3",
    srcs = [
        "detector.go",
        "s3.go",
    ],
    importpath = "go.resf.org/peridot/base/go/storage/s3",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//base/go/awsutils",
        "//base/go/storage",
        "//vendor/github.com/aws/aws-sdk-go/aws",
        "//vendor/github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor/github.com/aws/aws-sdk-go/aws/session",
        "//vendor/github.com/aws/aws-sdk-go/service/s3",
        "//vendor/github.com/aws/aws-sdk-go/service/s3/s3manager",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)
