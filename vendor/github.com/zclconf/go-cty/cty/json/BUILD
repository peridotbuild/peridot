load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "json",
    srcs = [
        "doc.go",
        "marshal.go",
        "simple.go",
        "type.go",
        "type_implied.go",
        "unmarshal.go",
        "value.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/zclconf/go-cty/cty/json",
    importpath = "github.com/zclconf/go-cty/cty/json",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/zclconf/go-cty/cty",
        "//vendor/github.com/zclconf/go-cty/cty/convert",
    ],
)
