load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "mathutil",
    srcs = [
        "binarylog.go",
        "bits.go",
        "envelope.go",
        "int.go",
        "mathutil.go",
        "permute.go",
        "poly.go",
        "primes.go",
        "rat.go",
        "rnd.go",
        "sqr.go",
        "sqr_std.go",
        "tables.go",
        "test_deps.go",
    ],
    importmap = "go.resf.org/peridot/vendor/modernc.org/mathutil",
    importpath = "modernc.org/mathutil",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:386": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:amd64": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:arm": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:arm64": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:mips": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:mips64": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:mips64le": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:mipsle": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:ppc64": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:ppc64le": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:s390x": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "@io_bazel_rules_go//go/platform:wasm": [
            "//vendor/github.com/remyoudompheng/bigfft",
        ],
        "//conditions:default": [],
    }),
)
