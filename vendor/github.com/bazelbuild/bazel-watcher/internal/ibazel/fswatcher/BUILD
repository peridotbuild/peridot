load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fswatcher",
    srcs = [
        "factory_darwin.go",
        "factory_fsnotify.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
    deps = select({
        "@io_bazel_rules_go//go/platform:aix": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:android": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsevents",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/log",
        ],
        "@io_bazel_rules_go//go/platform:dragonfly": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:freebsd": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:illumos": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:ios": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsevents",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/log",
        ],
        "@io_bazel_rules_go//go/platform:js": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:netbsd": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:openbsd": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:plan9": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:solaris": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
            "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/fsnotify",
        ],
        "//conditions:default": [],
    }),
)
