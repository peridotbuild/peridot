load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fxreflect",
    srcs = [
        "fxreflect.go",
        "stack.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.uber.org/fx/internal/fxreflect",
    importpath = "go.uber.org/fx/internal/fxreflect",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.uber.org/fx:__subpackages__",
    ],
)
