load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "handlers",
    srcs = [
        "canonical.go",
        "compress.go",
        "cors.go",
        "doc.go",
        "handlers.go",
        "logging.go",
        "proxy_headers.go",
        "recovery.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/gorilla/handlers",
    importpath = "github.com/gorilla/handlers",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/felixge/httpsnoop"],
)
