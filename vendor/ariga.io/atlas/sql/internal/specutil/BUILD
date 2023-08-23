load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "specutil",
    srcs = [
        "convert.go",
        "spec.go",
    ],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/internal/specutil",
    importpath = "ariga.io/atlas/sql/internal/specutil",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/ariga.io/atlas/sql:__subpackages__",
    ],
    deps = [
        "//vendor/ariga.io/atlas/schemahcl",
        "//vendor/ariga.io/atlas/sql/internal/sqlx",
        "//vendor/ariga.io/atlas/sql/schema",
        "//vendor/ariga.io/atlas/sql/sqlspec",
        "//vendor/github.com/hashicorp/hcl/v2/hclparse",
        "//vendor/github.com/zclconf/go-cty/cty",
    ],
)
