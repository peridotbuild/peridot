load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = ["internal.go"],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/xerrors/internal",
    importpath = "golang.org/x/xerrors/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/xerrors:__subpackages__",
    ],
)
