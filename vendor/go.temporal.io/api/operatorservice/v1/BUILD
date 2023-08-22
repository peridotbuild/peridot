load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "operatorservice",
    srcs = [
        "request_response.pb.go",
        "service.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/api/operatorservice/v1",
    importpath = "go.temporal.io/api/operatorservice/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/sortkeys",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
