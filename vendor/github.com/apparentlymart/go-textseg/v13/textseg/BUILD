load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "textseg",
    srcs = [
        "all_tokens.go",
        "generate.go",
        "grapheme_clusters.go",
        "tables.go",
        "utf8_seqs.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/apparentlymart/go-textseg/v13/textseg",
    importpath = "github.com/apparentlymart/go-textseg/v13/textseg",
    visibility = ["//visibility:public"],
)
