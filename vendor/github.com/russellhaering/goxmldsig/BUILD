load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "goxmldsig",
    srcs = [
        "canonicalize.go",
        "clock.go",
        "keystore.go",
        "sign.go",
        "tls_keystore.go",
        "validate.go",
        "xml_constants.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/russellhaering/goxmldsig",
    importpath = "github.com/russellhaering/goxmldsig",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/beevik/etree",
        "//vendor/github.com/jonboulle/clockwork",
        "//vendor/github.com/russellhaering/goxmldsig/etreeutils",
        "//vendor/github.com/russellhaering/goxmldsig/types",
    ],
)
