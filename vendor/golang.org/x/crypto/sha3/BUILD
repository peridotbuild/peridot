load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sha3",
    srcs = [
        "doc.go",
        "hashes.go",
        "hashes_generic.go",
        "keccakf.go",
        "keccakf_amd64.go",
        "keccakf_amd64.s",
        "register.go",
        "sha3.go",
        "sha3_s390x.go",
        "sha3_s390x.s",
        "shake.go",
        "shake_generic.go",
        "xor.go",
        "xor_generic.go",
        "xor_unaligned.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/crypto/sha3",
    importpath = "golang.org/x/crypto/sha3",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:s390x": [
            "//vendor/golang.org/x/sys/cpu",
        ],
        "//conditions:default": [],
    }),
)
