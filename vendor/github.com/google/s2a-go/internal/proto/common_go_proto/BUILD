load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "common_go_proto",
    srcs = ["common.pb.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/google/s2a-go/internal/proto/common_go_proto",
    importpath = "github.com/google/s2a-go/internal/proto/common_go_proto",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/google/s2a-go:__subpackages__",
    ],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
