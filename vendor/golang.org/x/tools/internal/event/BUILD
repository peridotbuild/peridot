load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "event",
    srcs = [
        "doc.go",
        "event.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/tools/internal/event",
    importpath = "golang.org/x/tools/internal/event",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/tools:__subpackages__",
    ],
    deps = [
        "//vendor/golang.org/x/tools/internal/event/core",
        "//vendor/golang.org/x/tools/internal/event/keys",
        "//vendor/golang.org/x/tools/internal/event/label",
    ],
)
