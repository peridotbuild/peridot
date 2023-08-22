load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "lib",
    srcs = [
        "ccgo.go",
        "cover.go",
        "etc.go",
        "go.go",
        "init.go",
        "mem.go",
        "mem_linux.go",
        "nodmesg.go",
        "stringer.go",
        "util.go",
    ],
    importmap = "go.resf.org/peridot/vendor/modernc.org/ccgo/v3/lib",
    importpath = "modernc.org/ccgo/v3/lib",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/kballard/go-shellquote",
        "//vendor/golang.org/x/tools/go/packages",
        "//vendor/modernc.org/cc/v3:cc",
        "//vendor/modernc.org/libc",
        "//vendor/modernc.org/mathutil",
        "//vendor/modernc.org/opt",
    ] + select({
        "@io_bazel_rules_go//go/platform:android": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "//conditions:default": [],
    }),
)
