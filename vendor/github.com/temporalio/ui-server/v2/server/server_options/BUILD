load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "server_options",
    srcs = [
        "apply_func_container.go",
        "server_option.go",
        "server_options.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/server_options",
    importpath = "github.com/temporalio/ui-server/v2/server/server_options",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/temporalio/ui-server/v2/server/api",
        "//vendor/github.com/temporalio/ui-server/v2/server/config",
    ],
)
