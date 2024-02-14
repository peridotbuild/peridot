load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "kernel_repack",
    srcs = ["v1.go"],
    embedsrcs = [
        "data/config-x86_64",
        "data/cpupower.config",
        "data/cpupower.service",
        "data/filter-modules.sh",
        "data/filter-x86_64.sh",
        "data/lt.spec",
        "data/ml.spec",
        "data/mod-extra.list",
        "data/config-aarch64",
        "data/filter-aarch64.sh",
        "data/kvm_stat.logrotate",
        "data/mod-denylist.sh",
        "data/mod-sign.sh",
        "data/secureboot-sig-kernel-x86_64.cer",
        "data/secureboot-sig-kernel-aarch64.cer",
        "data/rockykpatch1.x509",
        "data/x509.genkey",
        "data/rockydup1.x509",
    ],
    importpath = "go.resf.org/peridot/tools/kernelmanager/kernel_repack/v1",
    visibility = ["//visibility:public"],
    deps = ["//tools/kernelmanager/kernel_repack"],
)
