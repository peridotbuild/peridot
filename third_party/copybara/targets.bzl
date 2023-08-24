load("//tools/build_rules/copybara:copybara.bzl", "copybara")

copybara_list = [
    "googleapis",
    "bazel",
]

def declare_targets():
    for r in copybara_list:
        copybara(name = r)
