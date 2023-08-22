load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fxlog",
    srcs = [
        "default.go",
        "spy.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.uber.org/fx/internal/fxlog",
    importpath = "go.uber.org/fx/internal/fxlog",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.uber.org/fx:__subpackages__",
    ],
    deps = ["//vendor/go.uber.org/fx/fxevent"],
)
