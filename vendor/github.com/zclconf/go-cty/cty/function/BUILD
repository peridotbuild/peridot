load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "function",
    srcs = [
        "argument.go",
        "doc.go",
        "error.go",
        "function.go",
        "unpredictable.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/zclconf/go-cty/cty/function",
    importpath = "github.com/zclconf/go-cty/cty/function",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/zclconf/go-cty/cty"],
)
