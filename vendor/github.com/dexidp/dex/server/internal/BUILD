load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "codec.go",
        "types.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/server/internal",
    importpath = "github.com/dexidp/dex/server/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/dexidp/dex/server:__subpackages__",
    ],
    deps = [
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)
