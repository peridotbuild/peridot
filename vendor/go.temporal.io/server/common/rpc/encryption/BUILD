load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "encryption",
    srcs = [
        "localStoreCertProvider.go",
        "localStorePerHostCertProviderMap.go",
        "localStoreTlsProvider.go",
        "testDynamicCertProvider.go",
        "testDynamicTLSConfigProvider.go",
        "tlsFactory.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/rpc/encryption",
    importpath = "go.temporal.io/server/common/rpc/encryption",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/common/auth",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
    ],
)
