load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "testcontainersdocker",
    srcs = [
        "client.go",
        "docker_host.go",
        "docker_rootless.go",
        "docker_socket.go",
        "images.go",
        "labels.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/testcontainers/testcontainers-go/internal/testcontainersdocker",
    importpath = "github.com/testcontainers/testcontainers-go/internal/testcontainersdocker",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/testcontainers/testcontainers-go:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/docker/docker/client",
        "//vendor/github.com/testcontainers/testcontainers-go/internal/config",
        "//vendor/github.com/testcontainers/testcontainers-go/internal/testcontainerssession",
    ],
)
