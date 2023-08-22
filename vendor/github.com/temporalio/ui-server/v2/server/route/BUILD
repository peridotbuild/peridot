load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "route",
    srcs = [
        "api.go",
        "auth.go",
        "health.go",
        "openapi.go",
        "public_path.go",
        "ui.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/route",
    importpath = "github.com/temporalio/ui-server/v2/server/route",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-oidc/v3/oidc",
        "//vendor/github.com/gorilla/securecookie",
        "//vendor/github.com/labstack/echo/v4:echo",
        "//vendor/github.com/labstack/echo/v4/middleware",
        "//vendor/github.com/temporalio/ui-server/v2/server/api",
        "//vendor/github.com/temporalio/ui-server/v2/server/auth",
        "//vendor/github.com/temporalio/ui-server/v2/server/config",
        "@org_golang_x_net//context",
        "@org_golang_x_oauth2//:oauth2",
    ],
)
