load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "uuid",
    srcs = [
        "capi_darwin_amd64.go",
        "capi_darwin_arm64.go",
        "capi_linux_386.go",
        "capi_linux_amd64.go",
        "capi_linux_arm.go",
        "capi_linux_arm64.go",
        "capi_linux_ppc64le.go",
        "capi_linux_riscv64.go",
        "capi_linux_s390x.go",
        "uuid_darwin_amd64.go",
        "uuid_darwin_arm64.go",
        "uuid_linux_386.go",
        "uuid_linux_amd64.go",
        "uuid_linux_arm.go",
        "uuid_linux_arm64.go",
        "uuid_linux_ppc64le.go",
        "uuid_linux_riscv64.go",
        "uuid_linux_s390x.go",
    ],
    importmap = "go.resf.org/peridot/vendor/modernc.org/libc/uuid/uuid",
    importpath = "modernc.org/libc/uuid/uuid",
    visibility = ["//visibility:public"],
)
