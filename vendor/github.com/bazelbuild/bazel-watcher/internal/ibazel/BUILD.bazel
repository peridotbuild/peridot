load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ibazel",
    srcs = [
        "ibazel.go",
        "ibazel_unix.go",
        "ibazel_windows.go",
        "lifecycle.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
    deps = [
        "//third_party/github.com/bazelbuild/bazel-watcher/internal/ibazel/profiler",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/bazel",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/command",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/fswatcher/common",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/lifecycle_hooks",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/live_reload",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/log",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/output_runner",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/workspace",
        "//vendor/github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/blaze_query",
    ],
)
