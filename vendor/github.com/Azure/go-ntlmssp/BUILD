load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-ntlmssp",
    srcs = [
        "authenticate_message.go",
        "authheader.go",
        "avids.go",
        "challenge_message.go",
        "messageheader.go",
        "negotiate_flags.go",
        "negotiate_message.go",
        "negotiator.go",
        "nlmp.go",
        "unicode.go",
        "varfield.go",
        "version.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/Azure/go-ntlmssp",
    importpath = "github.com/Azure/go-ntlmssp",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/md4"],
)
