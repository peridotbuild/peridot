load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "pkgbits",
    srcs = [
        "codes.go",
        "decoder.go",
        "doc.go",
        "encoder.go",
        "flags.go",
        "frames_go1.go",
        "frames_go17.go",
        "reloc.go",
        "support.go",
        "sync.go",
        "syncmarker_string.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/tools/internal/pkgbits",
    importpath = "golang.org/x/tools/internal/pkgbits",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/tools:__subpackages__",
    ],
)
