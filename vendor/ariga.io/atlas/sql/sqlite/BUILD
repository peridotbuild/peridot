load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlite",
    srcs = [
        "convert.go",
        "diff.go",
        "driver.go",
        "driver_oss.go",
        "inspect.go",
        "migrate.go",
        "sqlspec.go",
    ],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/sqlite",
    importpath = "ariga.io/atlas/sql/sqlite",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/ariga.io/atlas/schemahcl",
        "//vendor/ariga.io/atlas/sql/internal/specutil",
        "//vendor/ariga.io/atlas/sql/internal/sqlx",
        "//vendor/ariga.io/atlas/sql/migrate",
        "//vendor/ariga.io/atlas/sql/schema",
        "//vendor/ariga.io/atlas/sql/sqlclient",
        "//vendor/ariga.io/atlas/sql/sqlspec",
        "//vendor/github.com/hashicorp/hcl/v2/hclparse",
        "//vendor/github.com/zclconf/go-cty/cty",
    ],
)
