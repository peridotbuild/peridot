load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cpuinfo",
    srcs = [
        "cpuinfo.go",
        "cpuinfo_amd64.go",
        "cpuinfo_amd64.s",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/klauspost/compress/internal/cpuinfo",
    importpath = "github.com/klauspost/compress/internal/cpuinfo",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/klauspost/compress:__subpackages__",
    ],
)
