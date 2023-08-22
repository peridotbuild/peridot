load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sdkrand",
    srcs = [
        "locked_source.go",
        "read.go",
        "read_1_5.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/aws/aws-sdk-go/internal/sdkrand",
    importpath = "github.com/aws/aws-sdk-go/internal/sdkrand",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/aws/aws-sdk-go:__subpackages__",
    ],
)
