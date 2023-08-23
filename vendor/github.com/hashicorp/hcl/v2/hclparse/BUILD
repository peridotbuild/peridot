load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "hclparse",
    srcs = ["parser.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/hashicorp/hcl/v2/hclparse",
    importpath = "github.com/hashicorp/hcl/v2/hclparse",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/hcl/v2:hcl",
        "//vendor/github.com/hashicorp/hcl/v2/hclsyntax",
        "//vendor/github.com/hashicorp/hcl/v2/json",
    ],
)
