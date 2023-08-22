load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "alias",
    srcs = ["alias.go"],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/crypto/internal/alias",
    importpath = "golang.org/x/crypto/internal/alias",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/crypto:__subpackages__",
    ],
)
