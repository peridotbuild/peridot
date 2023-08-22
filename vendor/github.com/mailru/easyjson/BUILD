load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "easyjson",
    srcs = [
        "helpers.go",
        "raw.go",
        "unknown_fields.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/mailru/easyjson",
    importpath = "github.com/mailru/easyjson",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mailru/easyjson/jlexer",
        "//vendor/github.com/mailru/easyjson/jwriter",
    ],
)
