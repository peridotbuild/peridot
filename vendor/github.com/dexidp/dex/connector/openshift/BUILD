load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "openshift",
    srcs = ["openshift.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/openshift",
    importpath = "github.com/dexidp/dex/connector/openshift",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/groups",
        "//vendor/github.com/dexidp/dex/pkg/httpclient",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/dexidp/dex/storage/kubernetes/k8sapi",
        "@org_golang_x_oauth2//:oauth2",
    ],
)
