load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "smetrics",
    srcs = [
        "doc.go",
        "hamming.go",
        "jaro.go",
        "jaro-winkler.go",
        "soundex.go",
        "ukkonen.go",
        "wagner-fischer.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/xrash/smetrics",
    importpath = "github.com/xrash/smetrics",
    visibility = ["//visibility:public"],
)
