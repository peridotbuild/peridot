load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "kernelorg",
    srcs = [
        "kernelorg.go",
        "releaseatom.go",
    ],
    embedsrcs = [
        "gregkh.asc",
        "torvalds.asc",
    ],
    importpath = "go.resf.org/peridot/tools/kernelmanager/kernel_repack/kernelorg",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/xi2/xz",
        "//vendor/golang.org/x/crypto/openpgp",
    ],
)
