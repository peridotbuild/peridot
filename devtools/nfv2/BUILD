load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "nfv2_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/devtools/nfv2",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//third_party/bazel/src/main/protobuf",
        "//vendor/golang.org/x/mod/modfile",
        "@org_golang_google_protobuf//proto",
    ],
)

go_binary(
    name = "nfv2",
    embed = [":nfv2_lib"],
    visibility = ["//visibility:public"],
)
