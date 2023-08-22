load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "s3err",
    srcs = ["error.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/aws/aws-sdk-go/internal/s3shared/s3err",
    importpath = "github.com/aws/aws-sdk-go/internal/s3shared/s3err",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/aws/aws-sdk-go:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor/github.com/aws/aws-sdk-go/aws/request",
    ],
)
