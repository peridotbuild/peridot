load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "utils",
    srcs = [
        "casing.go",
        "comments.go",
        "common_lints.go",
        "common_proto.go",
        "declarative_friendly.go",
        "extension.go",
        "find.go",
        "http.go",
        "message.go",
        "method.go",
        "resource.go",
        "string_pluralize.go",
        "type_name.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/rules/internal/utils",
    importpath = "github.com/googleapis/api-linter/rules/internal/utils",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/googleapis/api-linter/rules:__subpackages__",
    ],
    deps = [
        "//vendor/bitbucket.org/creachadair/stringset",
        "//vendor/cloud.google.com/go/longrunning/autogen/longrunningpb",
        "//vendor/github.com/gertd/go-pluralize",
        "//vendor/github.com/googleapis/api-linter/lint",
        "//vendor/github.com/googleapis/api-linter/locations",
        "//vendor/github.com/jhump/protoreflect/desc",
        "//vendor/github.com/jhump/protoreflect/desc/builder",
        "//vendor/github.com/stoewer/go-strcase",
        "@go_googleapis//google/api:annotations_go_proto",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)
