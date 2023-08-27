load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "armor",
    srcs = [
        "armor.go",
        "encode.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp/armor",
    importpath = "github.com/ProtonMail/go-crypto/openpgp/armor",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/ProtonMail/go-crypto/openpgp/errors"],
)
