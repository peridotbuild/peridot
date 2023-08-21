load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "parser",
    srcs = [
        "filter_base_listener.go",
        "filter_lexer.go",
        "filter_listener.go",
        "filter_parser.go",
        "gen.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.ciq.dev/pika/parser",
    importpath = "go.ciq.dev/pika/parser",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/antlr/antlr4/runtime/Go/antlr/v4:antlr"],
)
