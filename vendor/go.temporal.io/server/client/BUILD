load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "client",
    srcs = [
        "clientBean.go",
        "clientBean_mock.go",
        "clientFactory_mock.go",
        "clientfactory.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/client",
    importpath = "go.temporal.io/server/client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/server/api/adminservice/v1:adminservice",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/api/matchingservice/v1:matchingservice",
        "//vendor/go.temporal.io/server/client/admin",
        "//vendor/go.temporal.io/server/client/frontend",
        "//vendor/go.temporal.io/server/client/history",
        "//vendor/go.temporal.io/server/client/matching",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/membership",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/primitives",
    ],
)
