load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "copystructure",
    srcs = [
        "copier_time.go",
        "copystructure.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/mitchellh/copystructure",
    importpath = "github.com/mitchellh/copystructure",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/mitchellh/reflectwalk"],
)
