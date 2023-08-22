load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "compile",
    srcs = [
        "compile.go",
        "serial.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.starlark.net/internal/compile",
    importpath = "go.starlark.net/internal/compile",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.starlark.net:__subpackages__",
    ],
    deps = [
        "//vendor/go.starlark.net/resolve",
        "//vendor/go.starlark.net/syntax",
    ],
)
