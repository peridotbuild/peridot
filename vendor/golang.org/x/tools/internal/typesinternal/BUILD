load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "typesinternal",
    srcs = [
        "errorcode.go",
        "errorcode_string.go",
        "types.go",
        "types_118.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/tools/internal/typesinternal",
    importpath = "golang.org/x/tools/internal/typesinternal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/golang.org/x/tools:__subpackages__",
    ],
    deps = ["//vendor/golang.org/x/tools/go/types/objectpath"],
)
