load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "testhelper",
    srcs = [
        "certificate.go",
        "io.go",
        "source_root.go",
        "tls.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/tests/testhelper",
    importpath = "go.temporal.io/server/tests/testhelper",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/stretchr/testify/require"],
)
