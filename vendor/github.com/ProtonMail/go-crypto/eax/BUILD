load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "eax",
    srcs = [
        "eax.go",
        "eax_test_vectors.go",
        "random_vectors.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ProtonMail/go-crypto/eax",
    importpath = "github.com/ProtonMail/go-crypto/eax",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/ProtonMail/go-crypto/internal/byteutil"],
)
