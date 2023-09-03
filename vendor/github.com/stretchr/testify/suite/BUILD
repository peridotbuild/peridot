load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "suite",
    srcs = [
        "doc.go",
        "interfaces.go",
        "stats.go",
        "suite.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/stretchr/testify/suite",
    importpath = "github.com/stretchr/testify/suite",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/stretchr/testify/assert",
        "//vendor/github.com/stretchr/testify/require",
    ],
)
