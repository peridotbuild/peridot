load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "trace",
    srcs = ["trace.go"],
    importmap = "go.resf.org/peridot/vendor/cloud.google.com/go/internal/trace",
    importpath = "cloud.google.com/go/internal/trace",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/cloud.google.com/go:__subpackages__",
    ],
    deps = [
        "//vendor/go.opencensus.io/trace",
        "//vendor/golang.org/x/xerrors",
        "@go_googleapis//google/rpc:code_go_proto",
        "@org_golang_google_api//googleapi",
        "@org_golang_google_grpc//status",
    ],
)
