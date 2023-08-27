load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "hash",
    srcs = [
        "cyclic_poly.go",
        "doc.go",
        "rabin_karp.go",
        "roller.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/ulikunitz/xz/internal/hash",
    importpath = "github.com/ulikunitz/xz/internal/hash",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/ulikunitz/xz:__subpackages__",
    ],
)
