load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "locks",
    srcs = [
        "condition_variable.go",
        "condition_variable_impl.go",
        "id_mutex.go",
        "lock.go",
        "priority_mutex.go",
        "priority_mutex_impl.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/locks",
    importpath = "go.temporal.io/server/common/locks",
    visibility = ["//visibility:public"],
)
