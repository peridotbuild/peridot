load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlite",
    srcs = [
        "doc.go",
        "mutex.go",
        "norlimit.go",
        "rlimit_freebsd.go",
        "rlimit_ppc64le.go",
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
        "@io_bazel_rules_go//go/platform:freebsd_386": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:freebsd_amd64": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:freebsd_arm": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:freebsd_arm64": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:linux_ppc64le": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "//conditions:default": [],
    }),
)
