load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "packagesinternal",
    srcs = ["packages.go"],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/tools/internal/packagesinternal",
    importpath = "golang.org/x/tools/internal/packagesinternal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/tools:__subpackages__",
    ],
    deps = ["//vendor/golang.org/x/tools/internal/gocommand"],
)
