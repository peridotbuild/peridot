load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = ["version.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/testcontainers/testcontainers-go/internal",
    importpath = "github.com/testcontainers/testcontainers-go/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/testcontainers/testcontainers-go:__subpackages__",
    ],
)
