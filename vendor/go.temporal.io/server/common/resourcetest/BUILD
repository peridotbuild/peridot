load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "resourcetest",
    srcs = ["resourceTest.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/resourcetest",
    importpath = "go.temporal.io/server/common/resourcetest",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/uber-go/tally/v4:tally",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/api/workflowservicemock/v1:workflowservicemock",
        "//vendor/go.temporal.io/server/api/adminservice/v1:adminservice",
        "//vendor/go.temporal.io/server/api/adminservicemock/v1:adminservicemock",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/api/historyservicemock/v1:historyservicemock",
        "//vendor/go.temporal.io/server/api/matchingservice/v1:matchingservice",
        "//vendor/go.temporal.io/server/api/matchingservicemock/v1:matchingservicemock",
        "//vendor/go.temporal.io/server/client",
        "//vendor/go.temporal.io/server/common/archiver",
        "//vendor/go.temporal.io/server/common/archiver/provider",
        "//vendor/go.temporal.io/server/common/clock",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/membership",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/client",
        "//vendor/go.temporal.io/server/common/persistence/serialization",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/elasticsearch/client",
        "//vendor/go.temporal.io/server/common/primitives",
        "//vendor/go.temporal.io/server/common/sdk",
        "//vendor/go.temporal.io/server/common/searchattribute",
    ],
)
