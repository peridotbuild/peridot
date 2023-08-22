load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "standard",
    srcs = [
        "converter.go",
        "query_interceptors.go",
        "visibility_store.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/visibility/store/standard",
    importpath = "go.temporal.io/server/common/persistence/visibility/store/standard",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/xwb1989/sqlparser",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/common/convert",
        "//vendor/go.temporal.io/server/common/persistence/sql/sqlplugin",
        "//vendor/go.temporal.io/server/common/persistence/visibility/manager",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/elasticsearch",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/query",
        "//vendor/go.temporal.io/server/common/searchattribute",
    ],
)
