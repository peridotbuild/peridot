load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api",
    srcs = [
        "annotations.pb.go",
        "client.pb.go",
        "http.pb.go",
        "httpbody.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/gogo/googleapis/google/api",
    importpath = "github.com/gogo/googleapis/google/api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/descriptor",
        "//vendor/github.com/gogo/protobuf/types",
    ],
)
