load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "xerrors",
    srcs = [
        "adaptor.go",
        "doc.go",
        "errors.go",
        "fmt.go",
        "format.go",
        "frame.go",
        "wrap.go",
    ],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/xerrors",
    importpath = "golang.org/x/xerrors",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/xerrors/internal"],
)
