load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "uuid",
    srcs = [
        "capi_freebsd_386.go",
        "capi_freebsd_amd64.go",
        "capi_freebsd_arm.go",
        "capi_freebsd_arm64.go",
        "capi_netbsd_amd64.go",
        "capi_netbsd_arm.go",
        "capi_openbsd_386.go",
        "capi_openbsd_amd64.go",
        "capi_openbsd_arm64.go",
        "uuid_freebsd_386.go",
        "uuid_freebsd_amd64.go",
        "uuid_freebsd_arm.go",
        "uuid_freebsd_arm64.go",
        "uuid_netbsd_amd64.go",
        "uuid_netbsd_arm.go",
        "uuid_openbsd_386.go",
        "uuid_openbsd_amd64.go",
        "uuid_openbsd_arm64.go",
    ],
    importmap = "go.resf.org/peridot/vendor/modernc.org/libc/uuid",
    importpath = "modernc.org/libc/uuid",
    visibility = ["//visibility:public"],
)
