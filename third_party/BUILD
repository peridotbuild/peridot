exports_files(["openapi-generator-cli-6.6.0.jar"])

java_binary(
    name = "copybara_uberjar",
    main_class = "com.google.copybara.Main",
    runtime_deps = [":copybara_deploy.jar"],
)
