load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tally",
    srcs = [
        "histogram.go",
        "key_gen.go",
        "pool.go",
        "reporter.go",
        "sanitize.go",
        "scope.go",
        "scope_registry.go",
        "stats.go",
        "types.go",
        "version.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4",
    importpath = "github.com/uber-go/tally/v4",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/uber-go/tally/v4/internal/identity",
        "//vendor/go.uber.org/atomic",
    ],
)
