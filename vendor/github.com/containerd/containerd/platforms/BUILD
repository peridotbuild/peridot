load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "platforms",
    srcs = [
        "compare.go",
        "cpuinfo.go",
        "cpuinfo_linux.go",
        "cpuinfo_other.go",
        "database.go",
        "defaults.go",
        "defaults_darwin.go",
        "defaults_freebsd.go",
        "defaults_unix.go",
        "defaults_windows.go",
        "platforms.go",
        "platforms_other.go",
        "platforms_windows.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/containerd/containerd/platforms",
    importpath = "github.com/containerd/containerd/platforms",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/containerd/containerd/errdefs",
        "//vendor/github.com/containerd/containerd/log",
        "//vendor/github.com/opencontainers/image-spec/specs-go/v1:specs-go",
    ] + select({
        "@io_bazel_rules_go//go/platform:android": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/golang.org/x/sys/windows",
        ],
        "//conditions:default": [],
    }),
)
