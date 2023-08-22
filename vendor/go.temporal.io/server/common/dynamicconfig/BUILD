load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "dynamicconfig",
    srcs = [
        "client.go",
        "collection.go",
        "constants.go",
        "file_based_client.go",
        "file_based_client_mock.go",
        "shared_constants.go",
        "static_client.go",
        "static_properties.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/dynamicconfig",
    importpath = "go.temporal.io/server/common/dynamicconfig",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/primitives",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/gopkg.in/yaml.v3:yaml_v3",
    ],
)
