load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fsnotify",
    srcs = ["fsnotify.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
        "//vendor/github.com/fsnotify/fsnotify",
    ],
)
