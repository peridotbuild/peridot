load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ecdh",
    srcs = ["ecdh.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp/ecdh",
    importpath = "github.com/ProtonMail/go-crypto/openpgp/ecdh",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/aes/keywrap",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/internal/algorithm",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/internal/ecc",
    ],
)
