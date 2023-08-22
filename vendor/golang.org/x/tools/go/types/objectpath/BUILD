load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "objectpath",
    srcs = ["objectpath.go"],
    importmap = "go.resf.org/peridot/vendor/golang.org/x/tools/go/types/objectpath",
    importpath = "golang.org/x/tools/go/types/objectpath",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/tools/internal/typeparams"],
)
