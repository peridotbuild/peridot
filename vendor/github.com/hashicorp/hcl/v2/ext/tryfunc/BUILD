load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tryfunc",
    srcs = ["tryfunc.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/hashicorp/hcl/v2/ext/tryfunc",
    importpath = "github.com/hashicorp/hcl/v2/ext/tryfunc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/hcl/v2:hcl",
        "//vendor/github.com/hashicorp/hcl/v2/ext/customdecode",
        "//vendor/github.com/zclconf/go-cty/cty",
        "//vendor/github.com/zclconf/go-cty/cty/function",
    ],
)
