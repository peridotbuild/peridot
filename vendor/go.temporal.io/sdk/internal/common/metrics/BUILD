load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "metrics",
    srcs = [
        "capturing_handler.go",
        "constants.go",
        "grpc.go",
        "handler.go",
        "tags.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/sdk/internal/common/metrics",
    importpath = "go.temporal.io/sdk/internal/common/metrics",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/go.temporal.io/sdk:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/gogo/status",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/serviceerror",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
    ],
)
