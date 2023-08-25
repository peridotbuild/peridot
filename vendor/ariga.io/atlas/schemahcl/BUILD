load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "schemahcl",
    srcs = [
        "context.go",
        "extension.go",
        "schemahcl.go",
        "spec.go",
        "stdlib.go",
        "types.go",
    ],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/schemahcl",
    importpath = "ariga.io/atlas/schemahcl",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/ariga.io/atlas/sql/schema",
        "//vendor/github.com/go-openapi/inflect",
        "//vendor/github.com/hashicorp/hcl/v2:hcl",
        "//vendor/github.com/hashicorp/hcl/v2/ext/tryfunc",
        "//vendor/github.com/hashicorp/hcl/v2/gohcl",
        "//vendor/github.com/hashicorp/hcl/v2/hclparse",
        "//vendor/github.com/hashicorp/hcl/v2/hclsyntax",
        "//vendor/github.com/hashicorp/hcl/v2/hclwrite",
        "//vendor/github.com/zclconf/go-cty/cty",
        "//vendor/github.com/zclconf/go-cty/cty/convert",
        "//vendor/github.com/zclconf/go-cty/cty/function",
        "//vendor/github.com/zclconf/go-cty/cty/function/stdlib",
        "//vendor/github.com/zclconf/go-cty/cty/gocty",
        "//vendor/github.com/zclconf/go-cty/cty/json",
    ],
)
