load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "authpb",
    srcs = ["auth.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.etcd.io/etcd/api/v3/authpb",
    importpath = "go.etcd.io/etcd/api/v3/authpb",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)
