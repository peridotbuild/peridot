load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "serviceerror",
    srcs = [
        "convert.go",
        "currentBranchChanged.go",
        "retry_replication.go",
        "shardOwnershipLost.go",
        "sticky_worker_unavailable.go",
        "taskAlreadyStarted.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/serviceerror",
    importpath = "go.temporal.io/server/common/serviceerror",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/status",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/errordetails/v1:errordetails",
        "@org_golang_google_grpc//codes",
    ],
)
