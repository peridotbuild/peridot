load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bigfft",
    srcs = [
        "arith_decl.go",
        "fermat.go",
        "fft.go",
        "scan.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/remyoudompheng/bigfft",
    importpath = "github.com/remyoudompheng/bigfft",
    visibility = ["//visibility:public"],
)
