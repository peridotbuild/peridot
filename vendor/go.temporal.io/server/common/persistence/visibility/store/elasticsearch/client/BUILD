load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "client",
    srcs = [
        "aws.go",
        "bulk_processor.go",
        "bulk_processor_mock.go",
        "bulk_processor_v7.go",
        "bulk_service.go",
        "bulk_service_v7.go",
        "client.go",
        "client_factory.go",
        "client_mock.go",
        "client_v7.go",
        "config.go",
        "logger.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/visibility/store/elasticsearch/client",
    importpath = "go.temporal.io/server/common/persistence/visibility/store/elasticsearch/client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws",
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor/github.com/aws/aws-sdk-go/aws/session",
        "//vendor/github.com/blang/semver/v4:semver",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/olivere/elastic/v7:elastic",
        "//vendor/github.com/olivere/elastic/v7/aws/v4:aws",
        "//vendor/github.com/olivere/elastic/v7/uritemplates",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/common/log",
    ],
)
