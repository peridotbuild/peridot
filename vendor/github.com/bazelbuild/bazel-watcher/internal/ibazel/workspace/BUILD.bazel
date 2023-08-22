load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "workspace",
    srcs = ["workspace.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/workspace",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/workspace",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
    deps = ["//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/log"],
)
