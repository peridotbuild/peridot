load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ldap",
    srcs = ["ldap.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/ldap",
    importpath = "github.com/dexidp/dex/connector/ldap",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/go-ldap/ldap/v3:ldap",
    ],
)
