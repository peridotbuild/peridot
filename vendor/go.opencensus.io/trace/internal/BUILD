load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = ["internal.go"],
    importmap = "go.resf.org/peridot/vendor/go.opencensus.io/trace/internal",
    importpath = "go.opencensus.io/trace/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opencensus.io/trace:__subpackages__",
    ],
)
