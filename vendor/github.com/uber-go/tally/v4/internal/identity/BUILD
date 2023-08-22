load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "identity",
    srcs = ["accumulator.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/internal/identity",
    importpath = "github.com/uber-go/tally/v4/internal/identity",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/uber-go/tally/v4:__subpackages__",
    ],
    deps = ["//vendor/github.com/twmb/murmur3"],
)
