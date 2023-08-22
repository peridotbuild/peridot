load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "retry",
    srcs = ["interceptor.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/sdk/internal/common/retry",
    importpath = "go.temporal.io/sdk/internal/common/retry",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.temporal.io/sdk:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/gogo/status",
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware/retry",
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware/util/backoffutils",
        "//vendor/go.uber.org/atomic",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
    ],
)
