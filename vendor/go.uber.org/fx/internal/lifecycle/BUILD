load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "lifecycle",
    srcs = ["lifecycle.go"],
    importmap = "go.resf.org/peridot/vendor/go.uber.org/fx/internal/lifecycle",
    importpath = "go.uber.org/fx/internal/lifecycle",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.uber.org/fx:__subpackages__",
    ],
    deps = [
        "//vendor/go.uber.org/fx/fxevent",
        "//vendor/go.uber.org/fx/internal/fxclock",
        "//vendor/go.uber.org/fx/internal/fxreflect",
        "//vendor/go.uber.org/multierr",
    ],
)
