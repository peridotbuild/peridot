load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fsevents",
    srcs = [
        "fsevents.go",
        "go_1_10_after.go",
        "go_1_10_before.go",
        "go_1_9_2_before.go",
        "wrap.go",
    ],
    cgo = True,
    clinkopts = select({
        "@io_bazel_rules_go//go/platform:darwin": [
            "-framework CoreServices",
        ],
        "@io_bazel_rules_go//go/platform:ios": [
            "-framework CoreServices",
        ],
        "//conditions:default": [],
    }),
    importmap = "go.resf.org/peridot/vendor/github.com/fsnotify/fsevents",
    importpath = "github.com/fsnotify/fsevents",
    visibility = ["//visibility:public"],
)
