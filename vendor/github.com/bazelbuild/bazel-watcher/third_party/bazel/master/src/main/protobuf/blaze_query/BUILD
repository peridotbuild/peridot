load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "blaze_query",
    srcs = ["build.pb.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/blaze_query",
    importpath = "github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/blaze_query",
    visibility = ["//visibility:public"],
    deps = ["@com_github_golang_protobuf//proto:go_default_library"],
)
