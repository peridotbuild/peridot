load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "process_group",
    srcs = [
        "process_group.go",
        "process_group_unix.go",
        "process_group_windows.go",
        "syscalls_windows.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/process_group",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/process_group",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
)
