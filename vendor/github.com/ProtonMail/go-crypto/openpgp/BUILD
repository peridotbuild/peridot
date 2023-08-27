load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "openpgp",
    srcs = [
        "canonical_text.go",
        "hash.go",
        "key_generation.go",
        "keys.go",
        "keys_test_data.go",
        "read.go",
        "read_write_test_data.go",
        "write.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp",
    importpath = "github.com/ProtonMail/go-crypto/openpgp",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/armor",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/ecdh",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/ecdsa",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/eddsa",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/errors",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/internal/algorithm",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/internal/ecc",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/packet",
        "//vendor/golang.org/x/crypto/sha3",
    ],
)
