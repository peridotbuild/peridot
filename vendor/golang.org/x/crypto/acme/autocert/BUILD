load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "autocert",
    srcs = [
        "autocert.go",
        "cache.go",
        "listener.go",
        "renewal.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/crypto/acme/autocert",
    importpath = "golang.org/x/crypto/acme/autocert",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/crypto/acme",
        "@org_golang_x_net//idna",
    ],
)
