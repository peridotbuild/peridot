load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqltypes",
    srcs = [
        "bind_variables.go",
        "plan_value.go",
        "testing.go",
        "type.go",
        "value.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/xwb1989/sqlparser/dependency/sqltypes",
    importpath = "github.com/xwb1989/sqlparser/dependency/sqltypes",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/xwb1989/sqlparser/dependency/bytes2",
        "//vendor/github.com/xwb1989/sqlparser/dependency/hack",
        "//vendor/github.com/xwb1989/sqlparser/dependency/querypb",
    ],
)
