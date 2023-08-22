load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "annotate.go",
        "retry.go",
    ],
    importmap = "go.resf.org/peridot/vendor/cloud.google.com/go/internal",
    importpath = "cloud.google.com/go/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/cloud.google.com/go:__subpackages__",
    ],
    deps = [
        "@com_github_googleapis_gax_go_v2//:gax-go",
        "@org_golang_google_api//googleapi",
        "@org_golang_google_grpc//status",
    ],
)
