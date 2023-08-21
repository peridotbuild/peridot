load("//tools/build_rules/copybara:copybara.bzl", "copybara")

copybara_list = [
    "googleapis",
    "pika",
    "rules_dart",
]

def declare_targets():
    for r in copybara_list:
        copybara(name = r)
