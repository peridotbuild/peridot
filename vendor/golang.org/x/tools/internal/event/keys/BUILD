load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "keys",
    srcs = [
        "keys.go",
        "standard.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/tools/internal/event/keys",
    importpath = "golang.org/x/tools/internal/event/keys",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/tools:__subpackages__",
    ],
    deps = ["//vendor/golang.org/x/tools/internal/event/label"],
)
