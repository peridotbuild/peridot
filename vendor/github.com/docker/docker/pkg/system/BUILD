load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "system",
    srcs = [
        "args_windows.go",
        "chtimes.go",
        "chtimes_nowindows.go",
        "chtimes_windows.go",
        "errors.go",
        "filesys.go",
        "filesys_unix.go",
        "filesys_windows.go",
        "image_os.go",
        "init_windows.go",
        "lstat_unix.go",
        "lstat_windows.go",
        "meminfo_deprecated.go",
        "mknod.go",
        "mknod_freebsd.go",
        "mknod_unix.go",
        "mknod_windows.go",
        "path_deprecated.go",
        "process_deprecated.go",
        "stat_bsd.go",
        "stat_darwin.go",
        "stat_linux.go",
        "stat_openbsd.go",
        "stat_unix.go",
        "stat_windows.go",
        "utimes_unix.go",
        "utimes_unsupported.go",
        "xattrs_linux.go",
        "xattrs_unsupported.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/pkg/system",
    importpath = "github.com/docker/docker/pkg/system",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/meminfo",
    ] + select({
        "@io_bazel_rules_go//go/platform:aix": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:android": [
            "//vendor/github.com/docker/docker/pkg/process",
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "//vendor/github.com/docker/docker/pkg/process",
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:dragonfly": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:freebsd": [
            "//vendor/github.com/docker/docker/pkg/process",
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:illumos": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:ios": [
            "//vendor/github.com/docker/docker/pkg/process",
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:js": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/github.com/docker/docker/pkg/process",
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:netbsd": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:openbsd": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:plan9": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:solaris": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/github.com/docker/docker/pkg/process",
            "//vendor/golang.org/x/sys/windows",
        ],
        "//conditions:default": [],
    }),
)
