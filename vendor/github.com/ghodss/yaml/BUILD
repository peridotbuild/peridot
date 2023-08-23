load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "yaml",
    srcs = [
        "fields.go",
        "yaml.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ghodss/yaml",
    importpath = "github.com/ghodss/yaml",
    visibility = ["//visibility:public"],
    deps = ["//vendor/gopkg.in/yaml.v2:yaml_v2"],
)
