load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "echo",
    srcs = [
        "bind.go",
        "binder.go",
        "context.go",
        "context_fs.go",
        "echo.go",
        "echo_fs.go",
        "group.go",
        "group_fs.go",
        "ip.go",
        "json.go",
        "log.go",
        "response.go",
        "router.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/labstack/echo/v4",
    importpath = "github.com/labstack/echo/v4",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/labstack/gommon/color",
        "//vendor/github.com/labstack/gommon/log",
        "//vendor/golang.org/x/crypto/acme",
        "//vendor/golang.org/x/crypto/acme/autocert",
        "@org_golang_x_net//http2",
        "@org_golang_x_net//http2/h2c",
    ],
)
