load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "internal",
    srcs = [
        "message_context.go",
        "norace.go",
        "options.go",
        "race.go",
        "tags.go",
        "types.go",
        "util.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/bufbuild/protocompile/internal",
    importpath = "github.com/bufbuild/protocompile/internal",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/bufbuild/protocompile:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/bufbuild/protocompile/ast",
        "//vendor/github.com/bufbuild/protocompile/reporter",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
