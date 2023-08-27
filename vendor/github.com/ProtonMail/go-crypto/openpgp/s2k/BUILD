load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "s2k",
    srcs = [
        "s2k.go",
        "s2k_cache.go",
        "s2k_config.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp/s2k",
    importpath = "github.com/ProtonMail/go-crypto/openpgp/s2k",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/errors",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp/internal/algorithm",
        "//vendor/golang.org/x/crypto/argon2",
    ],
)
