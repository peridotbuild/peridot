load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "liteconfig",
    srcs = [
        "config.go",
        "freeport.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/temporalite/internal/liteconfig",
    importpath = "github.com/temporalio/temporalite/internal/liteconfig",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/temporalio/temporalite:__subpackages__",
    ],
    deps = [
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/persistence/sql/sqlplugin/sqlite",
        "//vendor/go.temporal.io/server/temporal",
    ],
)
