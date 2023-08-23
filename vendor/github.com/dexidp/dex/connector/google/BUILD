load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "google",
    srcs = ["google.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/google",
    importpath = "github.com/dexidp/dex/connector/google",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-oidc/v3/oidc",
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/groups",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/golang.org/x/exp/slices",
        "@org_golang_google_api//admin/directory/v1:directory",
        "@org_golang_google_api//option",
        "@org_golang_x_oauth2//:oauth2",
        "@org_golang_x_oauth2//google",
    ],
)
