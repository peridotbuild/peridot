load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "mothershippb_proto",
    srcs = [
        "bugtracker.proto",
        "entry.proto",
        "srpm_archiver.proto",
        "worker.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@com_google_protobuf//:wrappers_proto",
        "@go_googleapis//google/api:annotations_proto",
        "@googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "mothershippb_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
        "//:go_gen_grpc_gateway",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/pb",
    proto = ":mothershippb_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@go_googleapis//google/api:annotations_go_proto",
        "@org_golang_google_genproto//googleapis/api/annotations",
    ],
)

go_library(
    name = "pb",
    embed = [":mothershippb_go_proto"],
    importpath = "go.resf.org/peridot/tools/mothership/pb",
    visibility = ["//visibility:public"],
)
