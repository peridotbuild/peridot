load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "dex_lib",
    srcs = [
        "config.go",
        "main.go",
        "serve.go",
        "version.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/cmd/dex",
    importpath = "github.com/dexidp/dex/cmd/dex",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/AppsFlyer/go-sundheit",
        "//vendor/github.com/AppsFlyer/go-sundheit/checks",
        "//vendor/github.com/AppsFlyer/go-sundheit/http",
        "//vendor/github.com/dexidp/dex/api/v2:api",
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/dexidp/dex/server",
        "//vendor/github.com/dexidp/dex/storage",
        "//vendor/github.com/dexidp/dex/storage/ent",
        "//vendor/github.com/dexidp/dex/storage/etcd",
        "//vendor/github.com/dexidp/dex/storage/kubernetes",
        "//vendor/github.com/dexidp/dex/storage/memory",
        "//vendor/github.com/dexidp/dex/storage/sql",
        "//vendor/github.com/fsnotify/fsnotify",
        "//vendor/github.com/ghodss/yaml",
        "//vendor/github.com/grpc-ecosystem/go-grpc-prometheus",
        "//vendor/github.com/oklog/run",
        "//vendor/github.com/prometheus/client_golang/prometheus",
        "//vendor/github.com/prometheus/client_golang/prometheus/collectors",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
        "//vendor/github.com/sirupsen/logrus",
        "//vendor/github.com/spf13/cobra",
        "//vendor/golang.org/x/crypto/bcrypt",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials",
        "@org_golang_google_grpc//reflection",
    ],
)

go_binary(
    name = "dex",
    embed = [":dex_lib"],
    visibility = ["//visibility:public"],
)
