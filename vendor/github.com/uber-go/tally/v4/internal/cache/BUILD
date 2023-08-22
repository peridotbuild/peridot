load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cache",
    srcs = [
        "string_intern.go",
        "tag_cache.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/internal/cache",
    importpath = "github.com/uber-go/tally/v4/internal/cache",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/uber-go/tally/v4:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/uber-go/tally/v4/internal/identity",
        "//vendor/github.com/uber-go/tally/v4/m3/thrift/v2:thrift",
    ],
)
