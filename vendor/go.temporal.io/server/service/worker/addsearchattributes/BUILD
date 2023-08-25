load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "addsearchattributes",
    srcs = [
        "fx.go",
        "workflow.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/addsearchattributes",
    importpath = "go.temporal.io/server/service/worker/addsearchattributes",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/olivere/elastic/v7:elastic",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/sdk/temporal",
        "//vendor/go.temporal.io/sdk/worker",
        "//vendor/go.temporal.io/sdk/workflow",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/elasticsearch/client",
        "//vendor/go.temporal.io/server/common/searchattribute",
        "//vendor/go.temporal.io/server/common/util",
        "//vendor/go.temporal.io/server/service/worker/common",
        "//vendor/go.uber.org/fx",
        "//vendor/golang.org/x/exp/maps",
    ],
)
