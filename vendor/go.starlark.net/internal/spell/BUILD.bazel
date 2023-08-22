load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "spell",
    srcs = ["spell.go"],
    importmap = "go.resf.org/peridot/vendor/go.starlark.net/internal/spell",
    importpath = "go.starlark.net/internal/spell",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.starlark.net:__subpackages__",
    ],
)
