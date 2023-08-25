load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "runtime",
    srcs = [
        "gen.go",
        "plugin.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api/types/swarm/runtime",
    importpath = "github.com/docker/docker/api/types/swarm/runtime",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gogo/protobuf/proto"],
)
