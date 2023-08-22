load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "schema",
    srcs = [
        "version.go",
        "version_reader.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/schema",
    importpath = "go.temporal.io/server/common/persistence/schema",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/blang/semver/v4:semver"],
)
