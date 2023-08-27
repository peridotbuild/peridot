load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "encoding",
    srcs = [
        "encoding.go",
        "mpi.go",
        "oid.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/openpgp/internal/encoding",
    importpath = "github.com/ProtonMail/go-crypto/openpgp/internal/encoding",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/ProtonMail/go-crypto/openpgp:__subpackages__",
    ],
    deps = ["//vendor/github.com/ProtonMail/go-crypto/openpgp/errors"],
)
