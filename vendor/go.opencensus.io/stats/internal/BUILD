load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = ["record.go"],
    importmap = "go.resf.org/peridot/vendor/go.opencensus.io/stats/internal",
    importpath = "go.opencensus.io/stats/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opencensus.io/stats:__subpackages__",
    ],
    deps = ["//vendor/go.opencensus.io/tag"],
)
