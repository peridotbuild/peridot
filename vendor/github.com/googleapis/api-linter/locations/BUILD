load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "locations",
    srcs = [
        "descriptor_locations.go",
        "field_locations.go",
        "file_locations.go",
        "locations.go",
        "message_locations.go",
        "method_locations.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/googleapis/api-linter/locations",
    importpath = "github.com/googleapis/api-linter/locations",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/cloud.google.com/go/longrunning/autogen/longrunningpb",
        "//vendor/github.com/jhump/protoreflect/desc",
        "@go_googleapis//google/api:annotations_go_proto",
        "@io_bazel_rules_go//proto/wkt:descriptor_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
    ],
)
