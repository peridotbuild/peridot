load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "types",
    srcs = [
        "auth.go",
        "client.go",
        "configs.go",
        "error_response.go",
        "error_response_ext.go",
        "graph_driver_data.go",
        "id_response.go",
        "image_delete_response_item.go",
        "image_summary.go",
        "plugin.go",
        "plugin_device.go",
        "plugin_env.go",
        "plugin_interface_type.go",
        "plugin_mount.go",
        "plugin_responses.go",
        "port.go",
        "service_update_response.go",
        "stats.go",
        "types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api/types",
    importpath = "github.com/docker/docker/api/types",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types/container",
        "//vendor/github.com/docker/docker/api/types/filters",
        "//vendor/github.com/docker/docker/api/types/mount",
        "//vendor/github.com/docker/docker/api/types/network",
        "//vendor/github.com/docker/docker/api/types/registry",
        "//vendor/github.com/docker/docker/api/types/swarm",
        "//vendor/github.com/docker/docker/api/types/volume",
        "//vendor/github.com/docker/go-connections/nat",
        "//vendor/github.com/docker/go-units",
        "//vendor/github.com/opencontainers/image-spec/specs-go/v1:specs-go",
    ],
)
