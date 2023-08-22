load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "goro",
    srcs = [
        "goro.go",
        "group.go",
        "package.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/internal/goro",
    importpath = "go.temporal.io/server/internal/goro",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.temporal.io/server:__subpackages__",
    ],
)
