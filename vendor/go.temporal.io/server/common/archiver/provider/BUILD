load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "provider",
    srcs = [
        "provider.go",
        "provider_mock.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/archiver/provider",
    importpath = "go.temporal.io/server/common/archiver/provider",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/server/common/archiver",
        "//vendor/go.temporal.io/server/common/archiver/filestore",
        "//vendor/go.temporal.io/server/common/archiver/gcloud",
        "//vendor/go.temporal.io/server/common/archiver/s3store",
        "//vendor/go.temporal.io/server/common/config",
    ],
)
