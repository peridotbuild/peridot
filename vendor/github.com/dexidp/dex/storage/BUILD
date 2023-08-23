load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "storage",
    srcs = [
        "doc.go",
        "health.go",
        "static.go",
        "storage.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage",
    importpath = "github.com/dexidp/dex/storage",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/gopkg.in/square/go-jose.v2:go-jose_v2",
    ],
)
