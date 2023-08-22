load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "analysis",
    srcs = ["analysis.pb.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/analysis",
    importpath = "github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/analysis",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/blaze_query",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)
