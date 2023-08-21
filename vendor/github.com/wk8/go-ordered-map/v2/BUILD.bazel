load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-ordered-map",
    srcs = [
        "json.go",
        "orderedmap.go",
        "yaml.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/wk8/go-ordered-map/v2",
    importpath = "github.com/wk8/go-ordered-map/v2",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bahlo/generic-list-go",
        "//vendor/github.com/buger/jsonparser",
        "//vendor/github.com/mailru/easyjson/jwriter",
        "//vendor/gopkg.in/yaml.v3:yaml_v3",
    ],
)
