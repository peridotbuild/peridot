load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "standard_files.go",
        "unrecognized.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/jhump/protoreflect/internal",
    importpath = "github.com/jhump/protoreflect/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/jhump/protoreflect:__subpackages__",
    ],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
