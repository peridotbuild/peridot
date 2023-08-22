load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "shareddefaults",
    srcs = [
        "ecs_container.go",
        "shared_config.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/aws/aws-sdk-go/internal/shareddefaults",
    importpath = "github.com/aws/aws-sdk-go/internal/shareddefaults",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/aws/aws-sdk-go:__subpackages__",
    ],
)
