load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "aip0214",
    srcs = [
        "aip0214.go",
        "resource_expiry.go",
        "ttl_type.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/aip0214",
    importpath = "github.com/googleapis/api-linter/rules/aip0214",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/googleapis/api-linter/rules/internal/utils",
        "//vendor/github.com/jhump/protoreflect/desc",
    ],
)
