load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "github",
    srcs = ["github.go"],
    importpath = "go.resf.org/peridot/tools/mothership/worker_server/forge/github",
    visibility = ["//visibility:public"],
    deps = [
        "//tools/mothership/worker_server/forge",
        "//vendor/github.com/go-git/go-git/v5/plumbing/transport/http",
        "//vendor/github.com/golang-jwt/jwt/v5:jwt",
    ],
)
