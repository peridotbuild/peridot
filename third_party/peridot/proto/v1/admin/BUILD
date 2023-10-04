load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "adminpb_proto",
    srcs = ["admin.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "//third_party/peridot/proto/v1:peridotpb_proto",
        "@googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "adminpb_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
        "//:go_gen_grpc_gateway",
    ],
    importpath = "peridot.resf.org/peridot/admin/pb",
    proto = ":adminpb_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//third_party/peridot/proto/v1:pb",
        "@org_golang_google_genproto//googleapis/api/annotations",
    ],
)

go_library(
    name = "pb",
    embed = [":adminpb_go_proto"],
    importpath = "peridot.resf.org/peridot/admin/pb",
    visibility = ["//visibility:public"],
)
