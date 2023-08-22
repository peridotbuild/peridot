load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "ibazel_lib",
    srcs = [
        "main.go",
        "main_bsd.go",
        "main_linux.go",
        "main_unix.go",
        "main_windows.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bazelbuild/bazel-watcher/cmd/ibazel",
    importpath = "github.com/bazelbuild/bazel-watcher/cmd/ibazel",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel",
        "//vendor/github.com/bazelbuild/bazel-watcher/internal/ibazel/log",
    ] + select({
        "@io_bazel_rules_go//go/platform:android": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:dragonfly": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:freebsd": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:ios": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:netbsd": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:openbsd": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "//conditions:default": [],
    }),
)

go_binary(
    name = "ibazel",
    embed = [":ibazel_lib"],
    visibility = ["//visibility:public"],
)
