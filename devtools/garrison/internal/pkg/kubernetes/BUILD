load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "kubernetes",
    srcs = [
        "deployment.go",
        "env.go",
        "kubernetes.go",
        "metadata.go",
        "service.go",
        "virtual_service.go",
    ],
    importpath = "go.resf.org/peridot/devtools/garrison/internal/pkg/kubernetes",
    visibility = ["//devtools/garrison:__subpackages__"],
    deps = [
        "//base/go",
        "//devtools/garrison/internal/pkg/starlark",
        "//vendor/k8s.io/api/apps/v1:apps",
        "//vendor/k8s.io/api/core/v1:core",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/util/intstr",
        "//vendor/k8s.io/cli-runtime/pkg/printers",
    ],
)
