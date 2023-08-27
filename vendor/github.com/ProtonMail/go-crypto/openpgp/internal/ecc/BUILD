load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ecc",
    srcs = [
        "curve25519.go",
        "curve_info.go",
        "curves.go",
        "ed25519.go",
        "ed448.go",
        "generic.go",
        "x448.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp/internal/ecc",
    importpath = "github.com/ProtonMail/go-crypto/openpgp/internal/ecc",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/ProtonMail/go-crypto/bitcurves",
        "//vendor/github.com/ProtonMail/go-crypto/brainpool",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/errors",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/internal/encoding",
        "//vendor/github.com/cloudflare/circl/dh/x25519",
        "//vendor/github.com/cloudflare/circl/dh/x448",
        "//vendor/github.com/cloudflare/circl/sign/ed25519",
        "//vendor/github.com/cloudflare/circl/sign/ed448",
    ],
)
