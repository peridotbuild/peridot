load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "scrypt",
    srcs = ["scrypt.go"],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/crypto/scrypt",
    importpath = "golang.org/x/crypto/scrypt",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/pbkdf2"],
)
