load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "singleflight",
    srcs = ["singleflight.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/aws/aws-sdk-go/internal/sync/singleflight",
    importpath = "github.com/aws/aws-sdk-go/internal/sync/singleflight",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/aws/aws-sdk-go:__subpackages__",
    ],
)
