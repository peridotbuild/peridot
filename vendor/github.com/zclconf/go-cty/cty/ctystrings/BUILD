load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ctystrings",
    srcs = [
        "doc.go",
        "normalize.go",
        "prefix.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/zclconf/go-cty/cty/ctystrings",
    importpath = "github.com/zclconf/go-cty/cty/ctystrings",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/apparentlymart/go-textseg/v13/textseg",
        "//vendor/golang.org/x/text/unicode/norm",
    ],
)
