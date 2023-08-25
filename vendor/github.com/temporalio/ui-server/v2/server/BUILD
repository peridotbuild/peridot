load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "server",
    srcs = ["server.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server",
    importpath = "github.com/temporalio/ui-server/v2/server",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/labstack/echo/v4:echo",
        "//vendor/github.com/labstack/echo/v4/middleware",
        "//vendor/github.com/temporalio/ui-server/v2/openapi",
        "//vendor/github.com/temporalio/ui-server/v2/server/api",
        "//vendor/github.com/temporalio/ui-server/v2/server/auth",
        "//vendor/github.com/temporalio/ui-server/v2/server/config",
        "//vendor/github.com/temporalio/ui-server/v2/server/csrf",
        "//vendor/github.com/temporalio/ui-server/v2/server/headers",
        "//vendor/github.com/temporalio/ui-server/v2/server/route",
        "//vendor/github.com/temporalio/ui-server/v2/server/server_options",
        "//vendor/github.com/temporalio/ui-server/v2/ui",
    ],
)
