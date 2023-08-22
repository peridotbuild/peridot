load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = ["version.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/internal",
    importpath = "github.com/googleapis/api-linter/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/googleapis/api-linter:__subpackages__",
    ],
)
