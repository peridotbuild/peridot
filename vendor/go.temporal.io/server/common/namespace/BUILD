load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "namespace",
    srcs = [
        "archivalConfigStateMachine.go",
        "attrValidator.go",
        "const.go",
        "dlqMessageHandler.go",
        "dlqMessageHandler_mock.go",
        "errors.go",
        "fx.go",
        "handler.go",
        "handler_mock.go",
        "mutate.go",
        "namespace.go",
        "namespace_migration_util.go",
        "registry.go",
        "registry_mock.go",
        "replicationTaskExecutor.go",
        "replicationTaskHandler_mock.go",
        "testconstructors.go",
        "transmissionTaskHandler.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/namespace",
    importpath = "go.temporal.io/server/common/namespace",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/google/uuid",
        "//vendor/github.com/pborman/uuid",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/namespace/v1:namespace",
        "//vendor/go.temporal.io/api/replication/v1:replication",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/api/replication/v1:replication",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/archiver",
        "//vendor/go.temporal.io/server/common/archiver/provider",
        "//vendor/go.temporal.io/server/common/cache",
        "//vendor/go.temporal.io/server/common/clock",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/primitives",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/internal/goro",
        "//vendor/go.uber.org/fx",
        "//vendor/golang.org/x/exp/maps",
    ],
)
