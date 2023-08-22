load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "adminservicemock",
    srcs = ["service.pb.mock.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/adminservicemock/v1",
    importpath = "go.temporal.io/server/api/adminservicemock/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/server/api/adminservice/v1:adminservice",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
