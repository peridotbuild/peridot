load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "container",
    srcs = [
        "change_response_deprecated.go",
        "change_type.go",
        "change_types.go",
        "config.go",
        "container_top.go",
        "container_update.go",
        "create_response.go",
        "filesystem_change.go",
        "hostconfig.go",
        "hostconfig_unix.go",
        "hostconfig_windows.go",
        "wait_exit_error.go",
        "wait_response.go",
        "waitcondition.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api/types/container",
    importpath = "github.com/docker/docker/api/types/container",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types/blkiodev",
        "//vendor/github.com/docker/docker/api/types/mount",
        "//vendor/github.com/docker/docker/api/types/strslice",
        "//vendor/github.com/docker/go-connections/nat",
        "//vendor/github.com/docker/go-units",
    ],
)
