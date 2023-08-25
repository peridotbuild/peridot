load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "openapi",
    srcs = ["embed.go"],
    embedsrcs = [
        "assets/README.md",
        "assets/favicon-16x16.png",
        "assets/favicon-32x32.png",
        "assets/index.css",
        "assets/index.html",
        "assets/oauth2-redirect.html",
        "assets/swagger-initializer.js",
        "assets/swagger-ui-bundle.js",
        "assets/swagger-ui-bundle.js.map",
        "assets/swagger-ui-standalone-preset.js",
        "assets/swagger-ui-standalone-preset.js.map",
        "assets/swagger-ui.css",
        "assets/swagger-ui.css.map",
        "assets/swagger-ui.js",
        "assets/swagger-ui.js.map",
        "assets/temporal/api/workflowservice/v1/service.swagger.json",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/openapi",
    importpath = "github.com/temporalio/ui-server/v2/openapi",
    visibility = ["//visibility:public"],
)
