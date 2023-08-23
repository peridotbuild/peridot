load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cipher",
    srcs = [
        "cbc_hmac.go",
        "concat_kdf.go",
        "ecdh_es.go",
        "key_wrap.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/go-jose/go-jose/v3/cipher",
    importpath = "github.com/go-jose/go-jose/v3/cipher",
    visibility = ["//visibility:public"],
)
