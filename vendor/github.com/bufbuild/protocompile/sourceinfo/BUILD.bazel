load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sourceinfo",
    srcs = ["source_code_info.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/bufbuild/protocompile/sourceinfo",
    importpath = "github.com/bufbuild/protocompile/sourceinfo",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bufbuild/protocompile/ast",
        "//vendor/github.com/bufbuild/protocompile/internal",
        "//vendor/github.com/bufbuild/protocompile/options",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
