load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "data",
    srcs = ["data.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/internal/data",
    importpath = "github.com/googleapis/api-linter/rules/internal/data",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/googleapis/api-linter/rules:__subpackages__",
    ],
    deps = ["//vendor/bitbucket.org/creachadair/stringset"],
)
