def copybara(name):
    """Copybara is a tool for transforming and moving code between repositories."""
    native.sh_binary(
        name = name,
        srcs = ["//tools/build_rules/copybara:sync.bash"],
        env = {
            "TARGET_DIR": name,
        },
        visibility = ["//visibility:public"],
    )
