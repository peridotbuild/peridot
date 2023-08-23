load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bitbucketcloud",
    srcs = ["bitbucketcloud.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/bitbucketcloud",
    importpath = "github.com/dexidp/dex/connector/bitbucketcloud",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/groups",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "@org_golang_x_oauth2//:oauth2",
        "@org_golang_x_oauth2//bitbucket",
    ],
)
