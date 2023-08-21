load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0151",
    srcs = [
        "aip0151.go",
        "lro_metadata_reachable.go",
        "lro_metadata_type.go",
        "lro_response_reachable.go",
        "lro_response_type.go",
        "operation_info.go",
        "response_unary.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0151",
    importpath = "github.com/googleapis/api-linter/rules/aip0151",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
