load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "json",
    srcs = [
        "decode.go",
        "encode.go",
        "fold.go",
        "indent.go",
        "kubernetes_patch.go",
        "scanner.go",
        "stream.go",
        "tables.go",
        "tags.go",
    ],
    importmap = "go.resf.org/peridot/vendor/sigs.k8s.io/json/internal/golang/encoding/json",
    importpath = "sigs.k8s.io/json/internal/golang/encoding/json",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/sigs.k8s.io/json:__subpackages__",
    ],
)
