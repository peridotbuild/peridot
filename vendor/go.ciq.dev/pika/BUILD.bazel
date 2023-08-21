load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "pika",
    srcs = [
        "pika.go",
        "pika_aip_filter.go",
        "pika_aip_filter_proto.go",
        "pika_page_token.go",
        "pika_psql.go",
        "utils.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.ciq.dev/pika",
    importpath = "go.ciq.dev/pika",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/antlr/antlr4/runtime/Go/antlr/v4:antlr",
        "//vendor/github.com/gertd/go-pluralize",
        "//vendor/github.com/iancoleman/strcase",
        "//vendor/github.com/jmoiron/sqlx",
        "//vendor/github.com/lib/pq",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/sirupsen/logrus",
        "//vendor/github.com/wk8/go-ordered-map/v2:go-ordered-map",
        "//vendor/go.ciq.dev/pika/parser",
        "@org_golang_google_protobuf//encoding/protojson",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)
