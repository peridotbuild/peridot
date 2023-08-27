load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "memory",
    srcs = ["memory.go"],
    importpath = "go.resf.org/peridot/base/go/storage/memory",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go/storage",
        "//vendor/github.com/go-git/go-billy/v5:go-billy",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "memory_test",
    size = "small",
    srcs = ["memory_test.go"],
    embed = [":memory"],
    deps = [
        "//base/go/storage",
        "//vendor/github.com/go-git/go-billy/v5/memfs",
        "//vendor/github.com/stretchr/testify/require",
    ],
)
