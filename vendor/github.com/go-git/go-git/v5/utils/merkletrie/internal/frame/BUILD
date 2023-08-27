load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "frame",
    srcs = ["frame.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/go-git/go-git/v5/utils/merkletrie/internal/frame",
    importpath = "github.com/go-git/go-git/v5/utils/merkletrie/internal/frame",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/go-git/go-git/v5/utils/merkletrie:__subpackages__",
    ],
    deps = ["//vendor/github.com/go-git/go-git/v5/utils/merkletrie/noder"],
)
