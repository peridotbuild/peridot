load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "json",
    srcs = [
        "ast.go",
        "didyoumean.go",
        "doc.go",
        "is.go",
        "navigation.go",
        "parser.go",
        "peeker.go",
        "public.go",
        "scanner.go",
        "structure.go",
        "tokentype_string.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/hashicorp/hcl/v2/json",
    importpath = "github.com/hashicorp/hcl/v2/json",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/agext/levenshtein",
        "//vendor/github.com/apparentlymart/go-textseg/v13/textseg",
        "//vendor/github.com/hashicorp/hcl/v2:hcl",
        "//vendor/github.com/hashicorp/hcl/v2/hclsyntax",
        "//vendor/github.com/zclconf/go-cty/cty",
        "//vendor/github.com/zclconf/go-cty/cty/convert",
    ],
)
