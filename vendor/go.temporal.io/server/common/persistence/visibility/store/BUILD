load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "store",
    srcs = [
        "errors.go",
        "visibility_store.go",
        "visibility_store_mock.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/visibility/store",
    importpath = "go.temporal.io/server/common/persistence/visibility/store",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/visibility/manager",
    ],
)
