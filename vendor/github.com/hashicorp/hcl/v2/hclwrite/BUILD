load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "hclwrite",
    srcs = [
        "ast.go",
        "ast_attribute.go",
        "ast_block.go",
        "ast_body.go",
        "ast_expression.go",
        "doc.go",
        "format.go",
        "generate.go",
        "native_node_sorter.go",
        "node.go",
        "parser.go",
        "public.go",
        "tokens.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/hashicorp/hcl/v2/hclwrite",
    importpath = "github.com/hashicorp/hcl/v2/hclwrite",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/apparentlymart/go-textseg/v13/textseg",
        "//vendor/github.com/google/go-cmp/cmp",
        "//vendor/github.com/hashicorp/hcl/v2:hcl",
        "//vendor/github.com/hashicorp/hcl/v2/hclsyntax",
        "//vendor/github.com/zclconf/go-cty/cty",
    ],
)
