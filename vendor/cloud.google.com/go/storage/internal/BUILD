load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = ["version.go"],
    importmap = "go.resf.org/peridot/vendor/cloud.google.com/go/storage/internal",
    importpath = "cloud.google.com/go/storage/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/cloud.google.com/go/storage:__subpackages__",
    ],
)
