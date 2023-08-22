load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "primitives",
    srcs = [
        "namespaces.go",
        "role.go",
        "task_queues.go",
        "uuid.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/primitives",
    importpath = "go.temporal.io/server/common/primitives",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/uuid",
        "//vendor/github.com/pborman/uuid",
    ],
)
