load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "masker",
    srcs = ["masker.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/masker",
    importpath = "go.temporal.io/server/common/masker",
    visibility = ["//visibility:public"],
    deps = ["//vendor/gopkg.in/yaml.v3:yaml_v3"],
)
