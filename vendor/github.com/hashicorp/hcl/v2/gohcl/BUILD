load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gohcl",
    srcs = [
        "decode.go",
        "doc.go",
        "encode.go",
        "schema.go",
        "types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/hashicorp/hcl/v2/gohcl",
    importpath = "github.com/hashicorp/hcl/v2/gohcl",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/hcl/v2:hcl",
        "//vendor/github.com/hashicorp/hcl/v2/hclwrite",
        "//vendor/github.com/zclconf/go-cty/cty",
        "//vendor/github.com/zclconf/go-cty/cty/convert",
        "//vendor/github.com/zclconf/go-cty/cty/gocty",
    ],
)
