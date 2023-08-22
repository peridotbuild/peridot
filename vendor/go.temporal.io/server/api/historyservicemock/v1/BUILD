load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "historyservicemock",
    srcs = ["service.pb.mock.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/historyservicemock/v1",
    importpath = "go.temporal.io/server/api/historyservicemock/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
