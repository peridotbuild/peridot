load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "revision",
    srcs = [
        "parser.go",
        "scanner.go",
        "token.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/go-git/go-git/v5/internal/revision",
    importpath = "github.com/go-git/go-git/v5/internal/revision",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/go-git/go-git/v5:__subpackages__",
    ],
)
