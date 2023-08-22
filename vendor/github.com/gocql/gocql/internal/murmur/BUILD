load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "murmur",
    srcs = [
        "murmur.go",
        "murmur_appengine.go",
        "murmur_unsafe.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/gocql/gocql/internal/murmur",
    importpath = "github.com/gocql/gocql/internal/murmur",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/gocql/gocql:__subpackages__",
    ],
)
