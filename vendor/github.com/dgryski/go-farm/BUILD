load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-farm",
    srcs = [
        "basics.go",
        "farmhashcc.go",
        "farmhashmk.go",
        "farmhashna.go",
        "farmhashuo.go",
        "farmhashxo.go",
        "fp_amd64.s",
        "fp_generic.go",
        "fp_stub.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dgryski/go-farm",
    importpath = "github.com/dgryski/go-farm",
    visibility = ["//visibility:public"],
)
