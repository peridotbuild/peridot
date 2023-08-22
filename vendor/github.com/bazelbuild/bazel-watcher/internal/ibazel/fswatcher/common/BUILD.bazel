load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "common",
    srcs = ["common.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
    deps = ["//vendor/github.com/fsnotify/fsnotify"],
)
