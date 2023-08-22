load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "lint",
    srcs = [
        "config.go",
        "lint.go",
        "problem.go",
        "response.go",
        "rule.go",
        "rule_aliases.go",
        "rule_enabled.go",
        "rule_groups.go",
        "rule_name.go",
        "rule_registry.go",
        "rule_urls.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/lint",
    importpath = "github.com/googleapis/api-linter/lint",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bmatcuk/doublestar/v4:doublestar",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/gopkg.in/yaml.v2:yaml_v2",
        "@io_bazel_rules_go//proto/wkt:descriptor_go_proto",
    ],
)
