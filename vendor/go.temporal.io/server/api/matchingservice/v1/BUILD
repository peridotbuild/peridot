load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "matchingservice",
    srcs = [
        "request_response.pb.go",
        "service.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/matchingservice/v1",
    importpath = "go.temporal.io/server/api/matchingservice/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/sortkeys",
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/query/v1:query",
        "//vendor/go.temporal.io/api/taskqueue/v1:taskqueue",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/server/api/clock/v1:clock",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/history/v1:history",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
