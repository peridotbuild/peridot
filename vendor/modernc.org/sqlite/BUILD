load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlite",
    srcs = [
        "doc.go",
        "mutex.go",
        "norlimit.go",
        "rlimit.go",
        "rulimit.go",
        "sqlite.go",
        "sqlite_go18.go",
    ],
    importmap = "go.resf.org/peridot/vendor/modernc.org/sqlite",
    importpath = "modernc.org/sqlite",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/modernc.org/libc",
        "//vendor/modernc.org/libc/sys/types",
        "//vendor/modernc.org/sqlite/lib",
    ] + select({
        "@io_bazel_rules_go//go/platform:android": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
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
