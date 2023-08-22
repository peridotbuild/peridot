load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

# keep
go_library(
    name = "profiler",
    srcs = [
        "js.go",
        "profiler.go",
    ],
    importpath = "github.com/bazelbuild/bazel-watcher/internal/ibazel/profiler",  # keep
    visibility = ["//vendor/github.com/bazelbuild/bazel-watcher:__subpackages__"],  # keep
    deps = [
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/log",
        "//vendor/github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/blaze_query",
    ],
)
