load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sequential",
    srcs = [
        "doc.go",
        "sequential_unix.go",
        "sequential_windows.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/moby/sys/sequential",
    importpath = "github.com/moby/sys/sequential",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/golang.org/x/sys/windows",
        ],
        "//conditions:default": [],
    }),
)
