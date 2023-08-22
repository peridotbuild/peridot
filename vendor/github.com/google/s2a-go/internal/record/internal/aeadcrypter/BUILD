load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aeadcrypter",
    srcs = [
        "aeadcrypter.go",
        "aesgcm.go",
        "chachapoly.go",
        "common.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/google/s2a-go/internal/record/internal/aeadcrypter",
    importpath = "github.com/google/s2a-go/internal/record/internal/aeadcrypter",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/google/s2a-go:__subpackages__",
    ],
    deps = ["//vendor/golang.org/x/crypto/chacha20poly1305"],
)
