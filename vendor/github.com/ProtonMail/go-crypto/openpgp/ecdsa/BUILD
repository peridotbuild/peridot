load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ecdsa",
    srcs = ["ecdsa.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp/ecdsa",
    importpath = "github.com/ProtonMail/go-crypto/openpgp/ecdsa",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/ProtonMail/go-crypto/openpgp/internal/ecc"],
)
