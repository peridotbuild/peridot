load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = [
        "parse.go",
        "readfile.go",
        "sysreadfile.go",
        "sysreadfile_compat.go",
        "valueparser.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/prometheus/procfs/internal/util",
    importpath = "github.com/prometheus/procfs/internal/util",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/prometheus/procfs:__subpackages__",
    ],
)
