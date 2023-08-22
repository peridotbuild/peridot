load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "adminservice",
    srcs = [
        "request_response.pb.go",
        "service.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/adminservice/v1",
    importpath = "go.temporal.io/server/api/adminservice/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/sortkeys",
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/version/v1:version",
        "//vendor/go.temporal.io/api/workflow/v1:workflow",
        "//vendor/go.temporal.io/server/api/cluster/v1:cluster",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/history/v1:history",
        "//vendor/go.temporal.io/server/api/namespace/v1:namespace",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/api/replication/v1:replication",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
