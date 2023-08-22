load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@aspect_bazel_lib//lib:copy_to_bin.bzl", "copy_to_bin")

exports_files(["runner.sh"])

go_library(
    name = "taskrunner2_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/devtools/taskrunner2",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "taskrunner2",
    data = [
        "//vendor/github.com/bazelbuild/bazel-watcher/cmd/ibazel",
    ],
    embed = [":taskrunner2_lib"],
    visibility = ["//visibility:public"],
)
