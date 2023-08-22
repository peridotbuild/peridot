load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "manager",
    srcs = [
        "visibility_manager.go",
        "visibility_manager_mock.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/visibility/manager",
    importpath = "go.temporal.io/server/common/persistence/visibility/manager",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/workflow/v1:workflow",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
    ],
)
