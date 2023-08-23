load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "convert",
    srcs = [
        "compare_types.go",
        "conversion.go",
        "conversion_capsule.go",
        "conversion_collection.go",
        "conversion_dynamic.go",
        "conversion_object.go",
        "conversion_primitive.go",
        "conversion_tuple.go",
        "doc.go",
        "mismatch_msg.go",
        "public.go",
        "sort_types.go",
        "unify.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/zclconf/go-cty/cty/convert",
    importpath = "github.com/zclconf/go-cty/cty/convert",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/zclconf/go-cty/cty"],
)
