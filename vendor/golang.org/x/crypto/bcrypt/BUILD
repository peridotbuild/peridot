load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bcrypt",
    srcs = [
        "base64.go",
        "bcrypt.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/crypto/bcrypt",
    importpath = "golang.org/x/crypto/bcrypt",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/blowfish"],
)
