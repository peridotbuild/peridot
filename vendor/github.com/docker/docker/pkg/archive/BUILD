load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "archive",
    srcs = [
        "archive.go",
        "archive_linux.go",
        "archive_other.go",
        "archive_unix.go",
        "archive_windows.go",
        "changes.go",
        "changes_linux.go",
        "changes_other.go",
        "changes_unix.go",
        "changes_windows.go",
        "copy.go",
        "copy_unix.go",
        "copy_windows.go",
        "diff.go",
        "diff_unix.go",
        "diff_windows.go",
        "path.go",
        "path_unix.go",
        "path_windows.go",
        "time_linux.go",
        "time_unsupported.go",
        "whiteouts.go",
        "wrap.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/pkg/archive",
    importpath = "github.com/docker/docker/pkg/archive",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/containerd/containerd/pkg/userns",
        "//vendor/github.com/docker/docker/pkg/idtools",
        "//vendor/github.com/docker/docker/pkg/ioutils",
        "//vendor/github.com/docker/docker/pkg/pools",
        "//vendor/github.com/docker/docker/pkg/system",
        "//vendor/github.com/klauspost/compress/zstd",
        "//vendor/github.com/moby/patternmatcher",
        "//vendor/github.com/moby/sys/sequential",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/sirupsen/logrus",
    ] + select({
        "@io_bazel_rules_go//go/platform:aix": [
            "//vendor/golang.org/x/sys/unix",
        ],
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
        "@io_bazel_rules_go//go/platform:illumos": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:ios": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:js": [
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
        "@io_bazel_rules_go//go/platform:plan9": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:solaris": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/github.com/docker/docker/pkg/longpath",
        ],
        "//conditions:default": [],
    }),
)
