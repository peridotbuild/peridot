load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-grpc-middleware",
    srcs = [
        "doc.go",
        "wrappers.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/go-grpc-middleware/v2",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware/v2",
    visibility = ["//visibility:public"],
    deps = ["@org_golang_google_grpc//:go_default_library"],
)
