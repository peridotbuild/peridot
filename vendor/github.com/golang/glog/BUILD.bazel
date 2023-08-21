load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "glog",
    srcs = [
        "glog.go",
        "glog_file.go",
        "glog_flags.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/golang/glog",
    importpath = "github.com/golang/glog",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog/internal/logsink",
        "//vendor/github.com/golang/glog/internal/stackdump",
    ],
)
