load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "errdefs",
    srcs = [
        "errors.go",
        "grpc.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/containerd/containerd/errdefs",
    importpath = "github.com/containerd/containerd/errdefs",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
