load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "headers",
    srcs = [
        "caller_info.go",
        "headers.go",
        "versionChecker.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/headers",
    importpath = "go.temporal.io/server/common/headers",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/blang/semver/v4:semver",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/golang.org/x/exp/slices",
        "@org_golang_google_grpc//metadata",
    ],
)
