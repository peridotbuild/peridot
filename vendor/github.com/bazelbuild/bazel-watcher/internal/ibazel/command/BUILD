load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "command",
    srcs = [
        "command.go",
        "default_command.go",
        "notify_command.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/command",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/command",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/bazel",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/log",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/process_group",
    ],
)
