load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sockets",
    srcs = [
        "inmem_socket.go",
        "proxy.go",
        "sockets.go",
        "sockets_unix.go",
        "sockets_windows.go",
        "tcp_socket.go",
        "unix_socket.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/go-connections/sockets",
    importpath = "github.com/docker/go-connections/sockets",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_x_net//proxy",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/github.com/Microsoft/go-winio",
        ],
        "//conditions:default": [],
    }),
)
