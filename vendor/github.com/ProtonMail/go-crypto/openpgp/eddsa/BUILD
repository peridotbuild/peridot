load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "eddsa",
    srcs = ["eddsa.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp/eddsa",
    importpath = "github.com/ProtonMail/go-crypto/openpgp/eddsa",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/ProtonMail/go-crypto/openpgp/internal/ecc"],
)
