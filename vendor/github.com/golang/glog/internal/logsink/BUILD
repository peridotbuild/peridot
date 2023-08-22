load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "logsink",
    srcs = [
        "logsink.go",
        "logsink_fatal.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/golang/glog/internal/logsink",
    importpath = "github.com/golang/glog/internal/logsink",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/golang/glog:__subpackages__",
    ],
    deps = ["//vendor/github.com/golang/glog/internal/stackdump"],
)
