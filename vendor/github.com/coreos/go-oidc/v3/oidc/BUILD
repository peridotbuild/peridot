load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "oidc",
    srcs = [
        "jose.go",
        "jwks.go",
        "oidc.go",
        "verify.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/coreos/go-oidc/v3/oidc",
    importpath = "github.com/coreos/go-oidc/v3/oidc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/go-jose/go-jose/v3:go-jose",
        "@org_golang_x_oauth2//:oauth2",
    ],
)
