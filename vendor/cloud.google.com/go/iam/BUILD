load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "iam",
    srcs = ["iam.go"],
    importmap = "go.resf.org/peridot/vendor/cloud.google.com/go/iam",
    importpath = "cloud.google.com/go/iam",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/cloud.google.com/go/iam/apiv1/iampb",
        "@com_github_googleapis_gax_go_v2//:gax-go",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//metadata",
    ],
)
