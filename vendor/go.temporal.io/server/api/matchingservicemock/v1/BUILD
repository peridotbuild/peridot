load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "matchingservicemock",
    srcs = ["service.pb.mock.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/matchingservicemock/v1",
    importpath = "go.temporal.io/server/api/matchingservicemock/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/server/api/matchingservice/v1:matchingservice",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
