load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "csrf",
    srcs = ["skipper.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/csrf",
    importpath = "github.com/temporalio/ui-server/v2/server/csrf",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/labstack/echo/v4:echo"],
)
