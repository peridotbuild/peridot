load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "compact",
    srcs = [
        "compact.go",
        "language.go",
        "parents.go",
        "tables.go",
        "tags.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/text/internal/language/compact",
    importpath = "golang.org/x/text/internal/language/compact",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/text:__subpackages__",
    ],
    deps = ["//vendor/golang.org/x/text/internal/language"],
)
