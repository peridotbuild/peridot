load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "forge",
    srcs = [
        "caching.go",
        "forge.go",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/worker_server/forge",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/go-git/go-git/v5/plumbing/transport"],
)

go_test(
    name = "forge_test",
    size = "small",
    srcs = ["caching_test.go"],
    embed = [":forge"],
    deps = [
        "//vendor/github.com/go-git/go-git/v5/plumbing/transport/http",
        "//vendor/github.com/stretchr/testify/require",
    ],
)
