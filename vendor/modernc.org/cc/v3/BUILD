load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cc",
    srcs = [
        "abi.go",
        "abi_platforms.go",
        "ast.go",
        "ast2.go",
        "cc.go",
        "check.go",
        "cpp.go",
        "enum.go",
        "filesystem.go",
        "inspect.go",
        "lexer.go",
        "operand.go",
        "parser.go",
        "scanner.go",
        "stringer.go",
        "type.go",
    ],
    importmap = "go.resf.org/peridot/vendor/modernc.org/cc/v3",
    importpath = "modernc.org/cc/v3",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/lukechampine.com/uint128",
        "//vendor/modernc.org/mathutil",
        "//vendor/modernc.org/strutil",
        "//vendor/modernc.org/token",
    ],
)
