load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go",
    srcs = ["metrics.pb.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/prometheus/client_model/go",
    importpath = "github.com/prometheus/client_model/go",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)
