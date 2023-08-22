load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "lru",
    srcs = ["lru.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/gocql/gocql/internal/lru",
    importpath = "github.com/gocql/gocql/internal/lru",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/gocql/gocql:__subpackages__",
    ],
)
