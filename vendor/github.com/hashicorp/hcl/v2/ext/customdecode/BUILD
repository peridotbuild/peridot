load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "customdecode",
    srcs = [
        "customdecode.go",
        "expression_type.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/hashicorp/hcl/v2/ext/customdecode",
    importpath = "github.com/hashicorp/hcl/v2/ext/customdecode",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/hcl/v2:hcl",
        "//vendor/github.com/zclconf/go-cty/cty",
    ],
)
