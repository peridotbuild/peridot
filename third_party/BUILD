java_binary(
    name = "copybara_uberjar",
    main_class = "com.google.copybara.Main",
    runtime_deps = [":copybara_deploy.jar"],
)

java_binary(
    name = "openapi_generator",
    main_class = "org.openapitools.codegen.OpenAPIGenerator",
    visibility = ["//visibility:public"],
    runtime_deps = [":openapi-generator-cli-6.6.0.jar"],
)
