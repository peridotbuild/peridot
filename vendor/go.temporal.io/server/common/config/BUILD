load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "config",
    srcs = [
        "archival.go",
        "config.go",
        "loader.go",
        "localip.go",
        "persistence.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/config",
    importpath = "go.temporal.io/server/common/config",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gocql/gocql",
        "//vendor/go.temporal.io/server/common/auth",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/masker",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/persistence/visibility/store/elasticsearch/client",
        "//vendor/go.temporal.io/server/common/telemetry",
        "//vendor/gopkg.in/validator.v2:validator_v2",
        "//vendor/gopkg.in/yaml.v3:yaml_v3",
    ],
)
