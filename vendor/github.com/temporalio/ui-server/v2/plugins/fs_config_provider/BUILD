load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fs_config_provider",
    srcs = [
        "fs_config_provider.go",
        "loader.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/plugins/fs_config_provider",
    importpath = "github.com/temporalio/ui-server/v2/plugins/fs_config_provider",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/temporalio/ui-server/v2/server/config",
        "//vendor/gopkg.in/validator.v2:validator_v2",
        "//vendor/gopkg.in/yaml.v3:yaml_v3",
    ],
)
