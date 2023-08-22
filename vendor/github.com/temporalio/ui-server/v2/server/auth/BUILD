load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "auth",
    srcs = [
        "auth.go",
        "oidc.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/auth",
    importpath = "github.com/temporalio/ui-server/v2/server/auth",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-oidc/v3/oidc",
        "//vendor/github.com/labstack/echo/v4:echo",
        "//vendor/github.com/temporalio/ui-server/v2/server/config",
        "@org_golang_x_oauth2//:oauth2",
    ],
)
