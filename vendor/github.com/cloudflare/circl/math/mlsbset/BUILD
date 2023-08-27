load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "mlsbset",
    srcs = [
        "mlsbset.go",
        "power.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/cloudflare/circl/math/mlsbset",
    importpath = "github.com/cloudflare/circl/math/mlsbset",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/cloudflare/circl/internal/conv"],
)
