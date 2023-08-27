load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "armor",
    srcs = [
        "armor.go",
        "encode.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/crypto/openpgp/armor",
    importpath = "golang.org/x/crypto/openpgp/armor",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/openpgp/errors"],
)
