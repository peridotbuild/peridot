load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "schema",
    srcs = [
        "authcode.go",
        "authrequest.go",
        "client.go",
        "connector.go",
        "devicerequest.go",
        "devicetoken.go",
        "dialects.go",
        "keys.go",
        "offlinesession.go",
        "password.go",
        "refreshtoken.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/schema",
    importpath = "github.com/dexidp/dex/storage/ent/schema",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/entgo.io/ent",
        "//vendor/entgo.io/ent/dialect",
        "//vendor/entgo.io/ent/schema/field",
        "//vendor/github.com/dexidp/dex/storage",
        "//vendor/gopkg.in/square/go-jose.v2:go-jose_v2",
    ],
)
