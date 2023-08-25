load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "metadata",
    srcs = [
        "doc.go",
        "metadata.go",
        "single_key.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata",
    visibility = ["//visibility:public"],
    deps = ["@org_golang_google_grpc//metadata"],
)
