load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "auth",
    srcs = [
        "auth.go",
        "doc.go",
        "metadata.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware/v2:go-grpc-middleware",
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
