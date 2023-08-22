load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "codec",
    srcs = [
        "buffer.go",
        "decode.go",
        "encode.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/internal/codec",
    importpath = "github.com/jhump/protoreflect/internal/codec",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/jhump/protoreflect:__subpackages__",
    ],
    deps = ["@com_github_golang_protobuf//proto:go_default_library"],
)
