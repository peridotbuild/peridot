load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fs",
    srcs = ["fs.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/prometheus/procfs/internal/fs",
    importpath = "github.com/prometheus/procfs/internal/fs",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/prometheus/procfs:__subpackages__",
    ],
)
