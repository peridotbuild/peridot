load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "resolver",
    srcs = [
        "noop_resolver.go",
        "service_resolver.go",
        "service_resolver_mock.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/resolver",
    importpath = "go.temporal.io/server/common/resolver",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/mock/gomock"],
)
