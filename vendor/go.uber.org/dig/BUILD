load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "dig",
    srcs = [
        "constructor.go",
        "container.go",
        "cycle_error.go",
        "decorate.go",
        "doc.go",
        "error.go",
        "graph.go",
        "group.go",
        "inout.go",
        "invoke.go",
        "param.go",
        "provide.go",
        "result.go",
        "scope.go",
        "version.go",
        "visualize.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.uber.org/dig",
    importpath = "go.uber.org/dig",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.uber.org/dig/internal/digerror",
        "//vendor/go.uber.org/dig/internal/digreflect",
        "//vendor/go.uber.org/dig/internal/dot",
        "//vendor/go.uber.org/dig/internal/graph",
    ],
)
