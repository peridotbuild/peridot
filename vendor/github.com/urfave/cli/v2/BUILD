load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cli",
    srcs = [
        "app.go",
        "args.go",
        "category.go",
        "cli.go",
        "command.go",
        "context.go",
        "docs.go",
        "errors.go",
        "fish.go",
        "flag.go",
        "flag_bool.go",
        "flag_duration.go",
        "flag_ext.go",
        "flag_float64.go",
        "flag_float64_slice.go",
        "flag_generic.go",
        "flag_int.go",
        "flag_int64.go",
        "flag_int64_slice.go",
        "flag_int_slice.go",
        "flag_path.go",
        "flag_string.go",
        "flag_string_slice.go",
        "flag_timestamp.go",
        "flag_uint.go",
        "flag_uint64.go",
        "flag_uint64_slice.go",
        "flag_uint_slice.go",
        "funcs.go",
        "help.go",
        "parse.go",
        "sliceflag.go",
        "sort.go",
        "suggestions.go",
        "template.go",
        "zz_generated.flags.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/urfave/cli/v2",
    importpath = "github.com/urfave/cli/v2",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/cpuguy83/go-md2man/v2/md2man",
        "//vendor/github.com/xrash/smetrics",
    ],
)
