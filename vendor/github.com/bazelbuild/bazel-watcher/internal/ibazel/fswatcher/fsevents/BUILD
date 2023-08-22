load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fsevents",
    srcs = ["fsevents.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsevents",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsevents",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
    deps = select({
        "@io_bazel_rules_go//go/platform:darwin": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/fsnotify/fsevents",
        ],
        "@io_bazel_rules_go//go/platform:ios": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/fsnotify/fsevents",
        ],
        "//conditions:default": [],
    }),
)
