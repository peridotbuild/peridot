load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "output_runner",
    srcs = ["output_runner.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/output_runner",
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/output_runner",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/log",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/workspace",
        "//vendor/github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/blaze_query",
    ],
)
