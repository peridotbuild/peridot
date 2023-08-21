load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "walk",
    srcs = ["walk.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bufbuild/protocompile/walk",
    importpath = "github.com/bufbuild/protocompile/walk",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bufbuild/protocompile/internal",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
