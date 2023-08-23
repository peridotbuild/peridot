load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sqlx",
    srcs = [
        "dev.go",
        "diff.go",
        "exclude.go",
        "plan.go",
        "sqlx.go",
    ],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/internal/sqlx",
    importpath = "ariga.io/atlas/sql/internal/sqlx",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/ariga.io/atlas/sql:__subpackages__",
    ],
    deps = [
        "//vendor/ariga.io/atlas/sql/migrate",
        "//vendor/ariga.io/atlas/sql/schema",
    ],
)
