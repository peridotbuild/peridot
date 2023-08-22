load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "connector",
    srcs = [
        "client.go",
        "clientDelegate.go",
        "clientDelegate_mock.go",
        "client_mock.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/archiver/gcloud/connector",
    importpath = "go.temporal.io/server/common/archiver/gcloud/connector",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/cloud.google.com/go/storage",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/server/common/archiver",
        "//vendor/go.temporal.io/server/common/config",
        "@org_golang_google_api//iterator",
        "@org_golang_google_api//option",
        "@org_golang_x_oauth2//google",
    ],
)
