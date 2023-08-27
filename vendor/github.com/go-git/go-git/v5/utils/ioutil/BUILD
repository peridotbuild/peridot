load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ioutil",
    srcs = [
        "common.go",
        "pipe.go",
        "pipe_js.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/go-git/go-git/v5/utils/ioutil",
    importpath = "github.com/go-git/go-git/v5/utils/ioutil",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/jbenet/go-context/io",
    ] + select({
        "@io_bazel_rules_go//go/platform:js": [
            "//vendor/github.com/acomagu/bufpipe",
        ],
        "//conditions:default": [],
    }),
)
