load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-jose",
    srcs = [
        "asymmetric.go",
        "crypter.go",
        "doc.go",
        "encoding.go",
        "jwe.go",
        "jwk.go",
        "jws.go",
        "opaque.go",
        "shared.go",
        "signing.go",
        "symmetric.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/go-jose/go-jose/v3",
    importpath = "github.com/go-jose/go-jose/v3",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/go-jose/go-jose/v3/cipher",
        "//vendor/github.com/go-jose/go-jose/v3/json",
        "//vendor/golang.org/x/crypto/pbkdf2",
    ],
)
