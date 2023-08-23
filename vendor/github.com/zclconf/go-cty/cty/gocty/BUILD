load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gocty",
    srcs = [
        "doc.go",
        "helpers.go",
        "in.go",
        "out.go",
        "type_implied.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/zclconf/go-cty/cty/gocty",
    importpath = "github.com/zclconf/go-cty/cty/gocty",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/zclconf/go-cty/cty",
        "//vendor/github.com/zclconf/go-cty/cty/convert",
        "//vendor/github.com/zclconf/go-cty/cty/set",
    ],
)
