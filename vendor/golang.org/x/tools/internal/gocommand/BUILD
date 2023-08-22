load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gocommand",
    srcs = [
        "invoke.go",
        "vendor.go",
        "version.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/tools/internal/gocommand",
    importpath = "golang.org/x/tools/internal/gocommand",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/tools:__subpackages__",
    ],
    deps = [
        "//vendor/golang.org/x/mod/semver",
        "//vendor/golang.org/x/sys/execabs",
        "//vendor/golang.org/x/tools/internal/event",
    ],
)
