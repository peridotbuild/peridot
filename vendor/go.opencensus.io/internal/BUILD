load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "internal.go",
        "sanitize.go",
        "traceinternals.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.opencensus.io/internal",
    importpath = "go.opencensus.io/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.opencensus.io:__subpackages__",
    ],
    deps = ["//vendor/go.opencensus.io"],
)
