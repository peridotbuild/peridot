load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "meminfo",
    srcs = [
        "meminfo.go",
        "meminfo_linux.go",
        "meminfo_unsupported.go",
        "meminfo_windows.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/pkg/meminfo",
    importpath = "github.com/docker/docker/pkg/meminfo",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/golang.org/x/sys/windows",
        ],
        "//conditions:default": [],
    }),
)
