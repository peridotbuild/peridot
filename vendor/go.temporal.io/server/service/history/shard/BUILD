load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "shard",
    srcs = [
        "context.go",
        "context_impl.go",
        "context_mock.go",
        "context_testutil.go",
        "controller.go",
        "controller_impl.go",
        "controller_mock.go",
        "engine.go",
        "engine_factory.go",
        "engine_factory_mock.go",
        "engine_mock.go",
        "fx.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/shard",
    importpath = "go.temporal.io/server/service/history/shard",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.opentelemetry.io/otel/trace",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/history/v1:history",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/adminservice/v1:adminservice",
        "//vendor/go.temporal.io/server/api/clock/v1:clock",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/api/replication/v1:replication",
        "//vendor/go.temporal.io/server/client",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/archiver",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/clock",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/convert",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/future",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/membership",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/serialization",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/common/resourcetest",
        "//vendor/go.temporal.io/server/common/rpc",
        "//vendor/go.temporal.io/server/common/searchattribute",
        "//vendor/go.temporal.io/server/common/serviceerror",
        "//vendor/go.temporal.io/server/common/util",
        "//vendor/go.temporal.io/server/service/history/configs",
        "//vendor/go.temporal.io/server/service/history/consts",
        "//vendor/go.temporal.io/server/service/history/events",
        "//vendor/go.temporal.io/server/service/history/tasks",
        "//vendor/go.temporal.io/server/service/history/vclock",
        "//vendor/go.uber.org/fx",
        "//vendor/golang.org/x/exp/maps",
    ],
)
