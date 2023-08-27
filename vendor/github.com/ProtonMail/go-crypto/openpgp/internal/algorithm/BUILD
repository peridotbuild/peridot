load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "algorithm",
    srcs = [
        "aead.go",
        "cipher.go",
        "hash.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp/internal/algorithm",
    importpath = "github.com/ProtonMail/go-crypto/openpgp/internal/algorithm",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/ProtonMail/go-crypto/eax",
        "//vendor/github.com/ProtonMail/go-crypto/ocb",
        "//vendor/golang.org/x/crypto/cast5",
    ],
)
