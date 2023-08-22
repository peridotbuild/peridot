load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "service",
    srcs = ["service.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/google/s2a-go/internal/handshaker/service",
    importpath = "github.com/google/s2a-go/internal/handshaker/service",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/google/s2a-go:__subpackages__",
    ],
    deps = [
        "@org_golang_google_appengine//:appengine",
        "@org_golang_google_appengine//socket",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//grpclog",
    ],
)
