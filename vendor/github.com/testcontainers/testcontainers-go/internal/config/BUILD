load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "config",
    srcs = ["config.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/testcontainers/testcontainers-go/internal/config",
    importpath = "github.com/testcontainers/testcontainers-go/internal/config",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/testcontainers/testcontainers-go:__subpackages__",
    ],
    deps = ["//vendor/github.com/magiconair/properties"],
)
