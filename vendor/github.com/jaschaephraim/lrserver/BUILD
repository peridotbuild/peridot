load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "lrserver",
    srcs = [
        "connection.go",
        "handlers.go",
        "javascript.go",
        "lrserver.go",
        "messages.go",
        "server.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jaschaephraim/lrserver",
    importpath = "github.com/jaschaephraim/lrserver",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gorilla/websocket"],
)
