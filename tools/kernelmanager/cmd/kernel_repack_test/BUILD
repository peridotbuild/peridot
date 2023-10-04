load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "kernel_repack_test_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/kernelmanager/cmd/kernel_repack_test",
    visibility = ["//visibility:private"],
    deps = [
        "//tools/kernelmanager/kernel_repack/kernelorg",
        "//tools/kernelmanager/kernel_repack/v1:kernel_repack",
        "//vendor/github.com/go-git/go-billy/v5/osfs",
        "//vendor/github.com/pkg/errors",
    ],
)

go_binary(
    name = "kernel_repack_test",
    embed = [":kernel_repack_test_lib"],
    visibility = ["//visibility:public"],
)
