load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "dynamodb",
    srcs = ["dynamodb.go"],
    importpath = "go.resf.org/peridot/base/go/kv/dynamodb",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go/awsutils",
        "//base/go/kv",
        "//base/proto:pb",
        "//vendor/github.com/aws/aws-sdk-go/aws",
        "//vendor/github.com/aws/aws-sdk-go/aws/session",
        "//vendor/github.com/aws/aws-sdk-go/service/dynamodb",
        "@org_golang_google_protobuf//proto",
    ],
)
