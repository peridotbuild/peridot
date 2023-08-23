load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "client",
    srcs = [
        "authcode.go",
        "authrequest.go",
        "client.go",
        "connector.go",
        "devicerequest.go",
        "devicetoken.go",
        "keys.go",
        "main.go",
        "offlinesession.go",
        "password.go",
        "refreshtoken.go",
        "types.go",
        "utils.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/ent/client",
    importpath = "github.com/dexidp/dex/storage/ent/client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/storage",
        "//vendor/github.com/dexidp/dex/storage/ent/db",
        "//vendor/github.com/dexidp/dex/storage/ent/db/authcode",
        "//vendor/github.com/dexidp/dex/storage/ent/db/authrequest",
        "//vendor/github.com/dexidp/dex/storage/ent/db/devicerequest",
        "//vendor/github.com/dexidp/dex/storage/ent/db/devicetoken",
        "//vendor/github.com/dexidp/dex/storage/ent/db/migrate",
        "//vendor/github.com/dexidp/dex/storage/ent/db/password",
        "//vendor/github.com/pkg/errors",
    ],
)
