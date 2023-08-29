load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "error.go",
        "route_key.go",
        "stack_tracer.go",
        "submatches.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jarcoal/httpmock/internal",
    importpath = "github.com/jarcoal/httpmock/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/jarcoal/httpmock:__subpackages__",
    ],
)
