load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "wait",
    srcs = [
        "all.go",
        "errors.go",
        "errors_windows.go",
        "exec.go",
        "exit.go",
        "health.go",
        "host_port.go",
        "http.go",
        "log.go",
        "nop.go",
        "sql.go",
        "wait.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/testcontainers/testcontainers-go/wait",
    importpath = "github.com/testcontainers/testcontainers-go/wait",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types",
        "//vendor/github.com/docker/go-connections/nat",
        "//vendor/github.com/testcontainers/testcontainers-go/exec",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/golang.org/x/sys/windows",
        ],
        "//conditions:default": [],
    }),
)
