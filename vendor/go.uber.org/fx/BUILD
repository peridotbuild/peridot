load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fx",
    srcs = [
        "annotated.go",
        "app.go",
        "app_unixes.go",
        "app_windows.go",
        "decorate.go",
        "doc.go",
        "extract.go",
        "inout.go",
        "invoke.go",
        "lifecycle.go",
        "log.go",
        "module.go",
        "populate.go",
        "printer_writer.go",
        "provide.go",
        "replace.go",
        "shutdown.go",
        "supply.go",
        "version.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.uber.org/fx",
    importpath = "go.uber.org/fx",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.uber.org/dig",
        "//vendor/go.uber.org/fx/fxevent",
        "//vendor/go.uber.org/fx/internal/fxclock",
        "//vendor/go.uber.org/fx/internal/fxlog",
        "//vendor/go.uber.org/fx/internal/fxreflect",
        "//vendor/go.uber.org/fx/internal/lifecycle",
        "//vendor/go.uber.org/multierr",
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
        "@io_bazel_rules_go//go/platform:solaris": [
            "//vendor/golang.org/x/sys/unix",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/golang.org/x/sys/windows",
        ],
        "//conditions:default": [],
    }),
)
