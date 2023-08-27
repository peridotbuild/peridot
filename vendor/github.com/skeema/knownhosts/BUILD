load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "knownhosts",
    srcs = ["knownhosts.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/skeema/knownhosts",
    importpath = "github.com/skeema/knownhosts",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/crypto/ssh",
        "//vendor/golang.org/x/crypto/ssh/knownhosts",
    ],
)
