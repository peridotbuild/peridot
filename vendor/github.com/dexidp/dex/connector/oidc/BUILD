load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "oidc",
    srcs = ["oidc.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/oidc",
    importpath = "github.com/dexidp/dex/connector/oidc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-oidc/v3/oidc",
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/httpclient",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "@org_golang_x_oauth2//:oauth2",
    ],
)
