load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "asn1-ber",
    srcs = [
        "ber.go",
        "content_int.go",
        "generalizedTime.go",
        "header.go",
        "identifier.go",
        "length.go",
        "real.go",
        "util.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/go-asn1-ber/asn1-ber",
    importpath = "github.com/go-asn1-ber/asn1-ber",
    visibility = ["//visibility:public"],
)
