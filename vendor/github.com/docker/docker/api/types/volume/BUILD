load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "volume",
    srcs = [
        "cluster_volume.go",
        "create_options.go",
        "list_response.go",
        "options.go",
        "volume.go",
        "volume_update.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api/types/volume",
    importpath = "github.com/docker/docker/api/types/volume",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types/filters",
        "//vendor/github.com/docker/docker/api/types/swarm",
    ],
)
