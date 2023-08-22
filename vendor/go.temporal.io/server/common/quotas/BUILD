load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "quotas",
    srcs = [
        "dynamic.go",
        "dynamic_rate_limiter_impl.go",
        "multi_rate_limiter_impl.go",
        "multi_request_rate_limiter_impl.go",
        "multi_reservation_impl.go",
        "namespace_rate_limiter_impl.go",
        "noop_reservation_impl.go",
        "priority_rate_limiter_impl.go",
        "priority_reservation_impl.go",
        "rate_limiter.go",
        "rate_limiter_impl.go",
        "rate_limiter_mock.go",
        "request.go",
        "request_rate_limiter.go",
        "request_rate_limiter_mock.go",
        "reservation.go",
        "reservation_mock.go",
        "routing_rate_limiter_impl.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/quotas",
    importpath = "go.temporal.io/server/common/quotas",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.uber.org/atomic",
        "//vendor/golang.org/x/time/rate",
    ],
)
