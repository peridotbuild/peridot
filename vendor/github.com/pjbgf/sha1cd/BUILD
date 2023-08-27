load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sha1cd",
    srcs = [
        "detection.go",
        "sha1cd.go",
        "sha1cdblock_amd64.go",
        "sha1cdblock_amd64.s",
        "sha1cdblock_generic.go",
        "sha1cdblock_noasm.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/pjbgf/sha1cd",
    importpath = "github.com/pjbgf/sha1cd",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/pjbgf/sha1cd/internal",
        "//vendor/github.com/pjbgf/sha1cd/ubc",
    ],
)
