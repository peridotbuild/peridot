load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "remotesigner",
    srcs = ["remotesigner.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/google/s2a-go/internal/v2/remotesigner",
    importpath = "github.com/google/s2a-go/internal/v2/remotesigner",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/google/s2a-go:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/google/s2a-go/internal/proto/v2/s2a_go_proto",
        "//vendor/github.com/google/s2a-go/stream",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//grpclog",
    ],
)
