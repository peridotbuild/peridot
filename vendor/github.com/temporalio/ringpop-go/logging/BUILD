load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "logging",
    srcs = [
        "default.go",
        "facility.go",
        "level.go",
        "named.go",
        "nologger.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ringpop-go/logging",
    importpath = "github.com/temporalio/ringpop-go/logging",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/uber-common/bark"],
)
