load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "events",
    srcs = [
        "cache.go",
        "events_cache_mock.go",
        "notifier.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/events",
    importpath = "go.temporal.io/server/service/history/events",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/pborman/uuid",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/history/v1:history",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/cache",
        "//vendor/go.temporal.io/server/common/clock",
        "//vendor/go.temporal.io/server/common/collection",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
    ],
)
