load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "mysqlversion",
    srcs = ["mysqlversion.go"],
    embedsrcs = [
        "is/.README.md",
        "is/charset2collate",
        "is/charset2collate.maria",
        "is/collate2charset",
        "is/collate2charset.maria",
    ],
    importmap = "go.resf.org/peridot/vendor/ariga.io/atlas/sql/mysql/internal/mysqlversion",
    importpath = "ariga.io/atlas/sql/mysql/internal/mysqlversion",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/ariga.io/atlas/sql/mysql:__subpackages__",
    ],
    deps = [
        "//vendor/ariga.io/atlas/sql/internal/sqlx",
        "//vendor/ariga.io/atlas/sql/schema",
        "//vendor/golang.org/x/mod/semver",
    ],
)
