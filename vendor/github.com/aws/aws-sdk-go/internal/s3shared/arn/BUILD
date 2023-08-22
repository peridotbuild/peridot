load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "arn",
    srcs = [
        "accesspoint_arn.go",
        "arn.go",
        "outpost_arn.go",
        "s3_object_lambda_arn.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/aws/aws-sdk-go/internal/s3shared/arn",
    importpath = "github.com/aws/aws-sdk-go/internal/s3shared/arn",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/aws/aws-sdk-go:__subpackages__",
    ],
    deps = ["//vendor/github.com/aws/aws-sdk-go/aws/arn"],
)
