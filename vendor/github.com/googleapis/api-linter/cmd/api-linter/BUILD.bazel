load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "api-linter_lib",
    srcs = [
        "cli.go",
        "github_actions.go",
        "main.go",
        "rules.go",
        "summary.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/cmd/api-linter",
    importpath = "github.com/googleapis/api-linter/cmd/api-linter",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/googleapis/api-linter/internal",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/rules",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/desc/protoparse",
        "//vendor/github.com/olekukonko/tablewriter",
        "//vendor/github.com/spf13/pflag",
        "//vendor/gopkg.in/yaml.v2:yaml_v2",
        "@io_bazel_rules_go//proto/wkt:descriptor_go_proto",
        "@org_golang_google_protobuf//proto",
    ],
)

go_binary(
    name = "api-linter",
    embed = [":api-linter_lib"],
    visibility = ["//visibility:public"],
)
