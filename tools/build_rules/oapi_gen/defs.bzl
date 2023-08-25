load("@aspect_rules_js//js:defs.bzl", "js_library")
load("//third_party/grpc-gateway/protoc-gen-openapiv2:defs.bzl", "protoc_gen_openapiv2")
load("//tools/build_rules/ui_library:defs.bzl", "ui_library")
load(":ts.bzl", _oapi_gen_ts_sdk = "oapi_gen_ts_sdk")

oapi_gen_ts_sdk = _oapi_gen_ts_sdk

def oapi_gen_ts(name, proto, visibility = ["//visibility:private"], **kwargs):
    protoc_gen_openapiv2(
        name = name + "_openapiv2",
        proto = proto,
        simple_operation_ids = True,
        single_output = True,
        visibility = ["//visibility:private"],
    )

    oapi_gen_ts_sdk(
        name = name + "_gen",
        openapi = ":" + name + "_openapiv2",
        visibility = ["//visibility:private"],
        **kwargs
    )

    js_library(
        name = name,
        srcs = [name + "_gen"],
        visibility = visibility,
    )
