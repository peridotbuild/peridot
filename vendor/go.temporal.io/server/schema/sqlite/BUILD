load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlite",
    srcs = [
        "setup.go",
        "version.go",
    ],
    embedsrcs = [
        "v3/temporal/schema.sql",
        "v3/visibility/schema.sql",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/schema/sqlite",
    importpath = "go.temporal.io/server/schema/sqlite",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/serialization",
        "//vendor/go.temporal.io/server/common/persistence/sql",
        "//vendor/go.temporal.io/server/common/persistence/sql/sqlplugin",
        "//vendor/go.temporal.io/server/common/primitives",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/common/resolver",
    ],
)
