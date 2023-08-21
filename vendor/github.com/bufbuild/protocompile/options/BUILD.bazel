load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "options",
    srcs = ["options.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bufbuild/protocompile/options",
    importpath = "github.com/bufbuild/protocompile/options",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bufbuild/protocompile/ast",
        "//vendor/github.com/bufbuild/protocompile/internal",
        "//vendor/github.com/bufbuild/protocompile/linker",
        "//vendor/github.com/bufbuild/protocompile/parser",
        "//vendor/github.com/bufbuild/protocompile/reporter",
        "@org_golang_google_protobuf//encoding/protowire",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//types/descriptorpb",
        "@org_golang_google_protobuf//types/dynamicpb",
    ],
)
