load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlparser",
    srcs = [
        "analyzer.go",
        "ast.go",
        "comments.go",
        "encodable.go",
        "impossible_query.go",
        "normalizer.go",
        "parsed_query.go",
        "redact_query.go",
        "sql.go",
        "token.go",
        "tracked_buffer.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/xwb1989/sqlparser",
    importpath = "github.com/xwb1989/sqlparser",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/xwb1989/sqlparser/dependency/bytes2",
        "//vendor/github.com/xwb1989/sqlparser/dependency/querypb",
        "//vendor/github.com/xwb1989/sqlparser/dependency/sqltypes",
    ],
)
