load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "oauth",
    srcs = ["oauth.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/oauth",
    importpath = "github.com/dexidp/dex/connector/oauth",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/httpclient",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "@org_golang_x_oauth2//:oauth2",
    ],
)
