load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "saml",
    srcs = [
        "saml.go",
        "types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/saml",
    importpath = "github.com/dexidp/dex/connector/saml",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/beevik/etree",
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/groups",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/mattermost/xml-roundtrip-validator",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/russellhaering/goxmldsig",
        "//vendor/github.com/russellhaering/goxmldsig/etreeutils",
    ],
)
