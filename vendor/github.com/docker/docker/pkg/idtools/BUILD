load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "idtools",
    srcs = [
        "idtools.go",
        "idtools_unix.go",
        "idtools_windows.go",
        "usergroupadd_linux.go",
        "usergroupadd_unsupported.go",
        "utils_unix.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/pkg/idtools",
    importpath = "github.com/docker/docker/pkg/idtools",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:aix": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:android": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:dragonfly": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:freebsd": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:illumos": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:ios": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:js": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:netbsd": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:openbsd": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:plan9": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:solaris": [
            "//vendor/github.com/opencontainers/runc/libcontainer/user",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/github.com/docker/docker/pkg/system",
        ],
        "//conditions:default": [],
    }),
)
