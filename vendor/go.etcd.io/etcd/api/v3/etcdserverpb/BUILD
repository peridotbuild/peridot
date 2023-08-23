load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "etcdserverpb",
    srcs = [
        "etcdserver.pb.go",
        "raft_internal.pb.go",
        "raft_internal_stringer.go",
        "rpc.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.etcd.io/etcd/api/v3/etcdserverpb",
    importpath = "go.etcd.io/etcd/api/v3/etcdserverpb",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto",
        "//vendor/go.etcd.io/etcd/api/v3/authpb",
        "//vendor/go.etcd.io/etcd/api/v3/membershippb",
        "//vendor/go.etcd.io/etcd/api/v3/mvccpb",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@go_googleapis//google/api:annotations_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
