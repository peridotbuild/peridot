load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "worker",
    srcs = [
        "fx.go",
        "pernamespaceworker.go",
        "service.go",
        "worker.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker",
    importpath = "go.temporal.io/server/service/worker",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/sdk/client",
        "//vendor/go.temporal.io/sdk/worker",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/client",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/archiver",
        "//vendor/go.temporal.io/server/common/archiver/provider",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/membership",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/client",
        "//vendor/go.temporal.io/server/common/persistence/visibility",
        "//vendor/go.temporal.io/server/common/persistence/visibility/manager",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/elasticsearch/client",
        "//vendor/go.temporal.io/server/common/primitives",
        "//vendor/go.temporal.io/server/common/resolver",
        "//vendor/go.temporal.io/server/common/resource",
        "//vendor/go.temporal.io/server/common/sdk",
        "//vendor/go.temporal.io/server/common/searchattribute",
        "//vendor/go.temporal.io/server/service",
        "//vendor/go.temporal.io/server/service/worker/addsearchattributes",
        "//vendor/go.temporal.io/server/service/worker/archiver",
        "//vendor/go.temporal.io/server/service/worker/batcher",
        "//vendor/go.temporal.io/server/service/worker/common",
        "//vendor/go.temporal.io/server/service/worker/deletenamespace",
        "//vendor/go.temporal.io/server/service/worker/migration",
        "//vendor/go.temporal.io/server/service/worker/parentclosepolicy",
        "//vendor/go.temporal.io/server/service/worker/replicator",
        "//vendor/go.temporal.io/server/service/worker/scanner",
        "//vendor/go.temporal.io/server/service/worker/scheduler",
        "//vendor/go.uber.org/fx",
        "//vendor/golang.org/x/exp/maps",
    ],
)
