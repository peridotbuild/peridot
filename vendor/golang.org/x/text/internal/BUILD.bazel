load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "internal.go",
        "match.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/text/internal",
    importpath = "golang.org/x/text/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/text:__subpackages__",
    ],
    deps = ["//vendor/golang.org/x/text/language"],
)
