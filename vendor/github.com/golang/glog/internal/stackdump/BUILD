load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "stackdump",
    srcs = ["stackdump.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/golang/glog/internal/stackdump",
    importpath = "github.com/golang/glog/internal/stackdump",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/golang/glog:__subpackages__",
    ],
)
