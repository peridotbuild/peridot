load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ldap",
    srcs = [
        "add.go",
        "bind.go",
        "client.go",
        "compare.go",
        "conn.go",
        "control.go",
        "debug.go",
        "del.go",
        "dn.go",
        "doc.go",
        "error.go",
        "filter.go",
        "ldap.go",
        "moddn.go",
        "modify.go",
        "passwdmodify.go",
        "request.go",
        "search.go",
        "unbind.go",
        "whoami.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/go-ldap/ldap/v3",
    importpath = "github.com/go-ldap/ldap/v3",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-ntlmssp",
        "//vendor/github.com/go-asn1-ber/asn1-ber",
    ],
)
