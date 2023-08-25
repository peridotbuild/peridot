load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "testcontainerssession",
    srcs = ["session.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/testcontainers/testcontainers-go/internal/testcontainerssession",
    importpath = "github.com/testcontainers/testcontainers-go/internal/testcontainerssession",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/testcontainers/testcontainers-go:__subpackages__",
    ],
    deps = ["//vendor/github.com/google/uuid"],
)
